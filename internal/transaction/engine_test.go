package transaction

import (
	"context"
	"errors"
	"testing"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/adapter"
	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/capability"
)

func newTestEngine(t *testing.T, extra ...adapter.Adapter) (*Engine, *adapter.FileAdapter) { t.Helper(); fileAdapter,err:=adapter.NewFileAdapter(t.TempDir());if err!=nil{t.Fatal(err)};adapters:=append([]adapter.Adapter{fileAdapter},extra...);registry,err:=adapter.NewRegistry(adapters...);if err!=nil{t.Fatal(err)};store,err:=NewStore(t.TempDir());if err!=nil{t.Fatal(err)};engine,err:=NewEngine(registry,store);if err!=nil{t.Fatal(err)};return engine,fileAdapter }
func TestApplyAndRollback(t *testing.T){t.Parallel();engine,managed:=newTestEngine(t);ctx:=context.Background();plan:=Plan{Changes:[]Change{{Adapter:"local-file",Key:"theme/name",Action:ActionSet,Value:"dark"}}};preview,err:=engine.Preview(ctx,plan);if err!=nil{t.Fatal(err)};if !preview.Supported||preview.Changes[0].Noop{t.Fatalf("unexpected preview: %+v",preview)};tx,err:=engine.Apply(ctx,plan);if err!=nil{t.Fatal(err)};if tx.Status!=StatusApplied{t.Fatalf("status = %q",tx.Status)};got,exists,err:=managed.Read(ctx,"theme/name");if err!=nil||!exists||got!="dark"{t.Fatalf("managed state = (%q, %v, %v)",got,exists,err)};tx,err=engine.Rollback(ctx,tx.ID);if err!=nil{t.Fatal(err)};if tx.Status!=StatusRolledBack{t.Fatalf("rollback status = %q",tx.Status)};_,exists,err=managed.Read(ctx,"theme/name");if err!=nil{t.Fatal(err)};if exists{t.Fatal("managed state was not restored")};if _,err:=engine.Rollback(ctx,tx.ID);err!=nil{t.Fatalf("second rollback failed: %v",err)}}
func TestApplySameValueIsNoop(t *testing.T){t.Parallel();engine,managed:=newTestEngine(t);ctx:=context.Background();if err:=managed.Set(ctx,"theme/name","dark");err!=nil{t.Fatal(err)};plan:=Plan{Changes:[]Change{{Adapter:"local-file",Key:"theme/name",Action:ActionSet,Value:"dark"}}};tx,err:=engine.Apply(ctx,plan);if err!=nil{t.Fatal(err)};if !tx.Changes[0].Noop{t.Fatal("same desired value should be a no-op")}}
func TestUnsupportedAdapterIsReportedBeforeMutation(t *testing.T){t.Parallel();engine,_:=newTestEngine(t);ctx:=context.Background();plan:=Plan{Changes:[]Change{{Adapter:"does-not-exist",Key:"anything",Action:ActionSet,Value:"x"}}};preview,err:=engine.Preview(ctx,plan);if err!=nil{t.Fatal(err)};if preview.Supported||preview.Changes[0].Supported{t.Fatalf("unsupported adapter unexpectedly supported: %+v",preview)};if _,err:=engine.Apply(ctx,plan);!errors.Is(err,ErrUnsupported){t.Fatalf("apply error = %v",err)}}
func TestPartialFailureRollsBackEarlierChanges(t *testing.T){t.Parallel();fail:=&failingAdapter{};engine,managed:=newTestEngine(t,fail);ctx:=context.Background();plan:=Plan{Changes:[]Change{{Adapter:"local-file",Key:"theme/name",Action:ActionSet,Value:"dark"},{Adapter:"fail",Key:"boom",Action:ActionSet,Value:"x"}}};tx,err:=engine.Apply(ctx,plan);if err==nil{t.Fatal("expected apply failure")};if tx.Status!=StatusRolledBack{t.Fatalf("status = %q; err=%v",tx.Status,err)};_,exists,readErr:=managed.Read(ctx,"theme/name");if readErr!=nil{t.Fatal(readErr)};if exists{t.Fatal("first change survived a later apply failure")}}

func TestValidationMismatchRollsBackJustWrittenChange(t *testing.T){
	t.Parallel(); a:=&lyingAdapter{value:"before",exists:true}; engine,_:=newTestEngine(t,a); ctx:=context.Background()
	plan:=Plan{Changes:[]Change{{Adapter:"lying",Key:"setting",Action:ActionSet,Value:"after"}}}
	tx,err:=engine.Apply(ctx,plan);if err==nil{t.Fatal("expected validation failure")};if tx.Status!=StatusRolledBack{t.Fatalf("status = %q, want rolled_back; err=%v",tx.Status,err)};if a.value!="before"||!a.exists{t.Fatalf("validation failure did not restore snapshot: value=%q exists=%v",a.value,a.exists)}
}
func TestValidationReadFailureRollsBackJustWrittenChange(t *testing.T){
	t.Parallel(); a:=&lyingAdapter{value:"before",exists:true,failReadAfterSet:true}; engine,_:=newTestEngine(t,a); ctx:=context.Background()
	plan:=Plan{Changes:[]Change{{Adapter:"lying",Key:"setting",Action:ActionSet,Value:"after"}}}
	tx,err:=engine.Apply(ctx,plan);if err==nil{t.Fatal("expected validation read failure")};if tx.Status!=StatusRolledBack{t.Fatalf("status = %q; err=%v",tx.Status,err)};if a.value!="before"||!a.exists{t.Fatalf("read failure did not restore snapshot: value=%q exists=%v",a.value,a.exists)}
}

type failingAdapter struct{}
func(*failingAdapter)Name()string{return"fail"};func(*failingAdapter)Capabilities(context.Context)[]capability.Capability{return[]capability.Capability{{ID:"test-failure",Supported:true}}};func(*failingAdapter)Read(context.Context,string)(string,bool,error){return"",false,nil};func(*failingAdapter)Set(context.Context,string,string)error{return errors.New("injected failure")};func(*failingAdapter)Unset(context.Context,string)error{return errors.New("injected failure")}

type lyingAdapter struct{value string;exists bool;set bool;failReadAfterSet bool}
func(*lyingAdapter)Name()string{return"lying"}
func(*lyingAdapter)Capabilities(context.Context)[]capability.Capability{return[]capability.Capability{{ID:"test-validation",Supported:true}}}
func(a *lyingAdapter)Read(context.Context,string)(string,bool,error){if a.set&&a.failReadAfterSet{return"",false,errors.New("injected validation read failure")};return a.value,a.exists,nil}
func(a *lyingAdapter)Set(_ context.Context,_ string,value string)error{if a.set { a.value=value;a.exists=true;return nil };a.set=true;a.value="wrong";a.exists=true;return nil}
func(a *lyingAdapter)Unset(context.Context,string)error{a.exists=false;a.value="";return nil}
