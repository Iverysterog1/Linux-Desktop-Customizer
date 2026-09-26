#!/usr/bin/env bash
set -Eeuo pipefail

readonly EXPECTED_SHA256='c1e95430c5fe2bba0259216240c8d056686f74ba09b9f1607ea5d8e33af6fce5'
readonly EXPECTED_VERSION='0.9.0-beta.1'
readonly WORK_DIR="${WORK_DIR:-$(mktemp -d)}"
readonly ARCHIVE="$WORK_DIR/ldc-beta.tar.xz"
readonly SOURCE="$WORK_DIR/source"

cleanup() { [[ "${KEEP_WORK_DIR:-0}" == 1 ]] || rm -rf "$WORK_DIR"; }
trap cleanup EXIT

: "${GITHUB_REPOSITORY:?GITHUB_REPOSITORY is required}"
: "${GH_TOKEN:?GH_TOKEN is required}"
for command in gh base64 sha256sum xz tar go; do command -v "$command" >/dev/null; done

blobs=(
617acebce23f59124269e03c65f0fd6bb3b1be70 cf03a0e7202be9a3f0b7897a704b4087c9814719
673858ea8028133c68dd700c824cf22d610184a4 f8042331f3cfa662a93afdd04a7231703dd98075
cfd027fdf53fcd403a020f7a934b74d0febdc98c 3695085d2f0a69956dd155d465e2725942b48385
351ff955e694f658a8f6f39b2458503a70bbcf78 e94ad44b6225f3b83d8e36d46560beea27264772
bb33872518f7b9d18e21ac385f76676fe78cf36b 86e8fc45d980475d88382e534f30c98eef1af150
14afbd78b9c189819508710e1f38db87f31585a1 68ce7aca4a1189d991efa756e6235fee797adf12
e726f3d796111698402c939d0d1084a408539601 9d1eda00dddc4789b5a5cfede274e3f4baabd521
84098dd1fb962cb322892484152af3972613447b 611ddeebe7f37a630532ffe0b6a375987d773ce5
f9527a156e649820dcc3826afaf621f592356ffc 2a7ebd80064c5962588ce59ad45d38e51a0fbd2b
feee19d94448edac055434426aea8108dacf32fa da6d0b672526f460019d92732def7fd5c5d77752
056b41787ab7f789e08ea9bd370de26096988b81 918d2d70b6f552339bfc2b13c96f7a10142a6a6e
f8c52047d068099e45a6ef3b611435cc81f92b6d 6f6723095bb5f0c5c3fad27b5116ca3a3337b5ef
92ed0eb683bcfd64708c7f3d9fa6f428d18af4c8 4fbb923e7f3178b22445ebf30f408281911f5d46
7c4ec5ecc6ae93bc90444a4a5da45a8f0bb00205 09645ef3ca9f3564395ddf762f9f8f83410507d2
7f2cb8157cb834371d5c2c6dbdd5974e95360ada 762d023ca08c022cbb1a912030766f6972fdc089
a667652441125582e387057999c3fa0d759b0ad1 d6f0c666cd823588205677566986f4d3132ce7dd
db70ff24edf403409a59c00d7107a963800d080b e7996255197a029cb24076920057f45fc9da771c
0251a93e1bb2b651802868a01321e222ecf6a05e 8943b96af3f1109965494226093441dffb3d39c6
c86ca53447c723482314cedebe1999898c481351 756ec8cc03167083ef3479462416814962376c8b
c3ee3ae69cbda22af78d3c5f5d3f147938a3e8cd b951123c696562f43d36387189b2dd1fbb19c080
33d3302ec4ab5f5af23ce85f9a85cddc803ff225 51ed7217500bbfafcf4fce1acc33ac4ce92f7718
9c3b2f0e2282906291b0ff49f42f4b7bab5cf658 18c415334fb1bc0e22d5fff6d109b3bd7da30ca7
14e0587dc7744354b269368388730300ff69efc2 9db25ade812e750f4addec256cb7ca8fa4d14f61
4d8cb5bc0d391289fdc1ff00b9a52d71bb3ec913 06f30797eb51c9f8da76186cf33a43ec430e1494
c16bf3f436126eb0fb3db84e341939cc57cf26da 1ef76892751a89571cdf00c47892a1f32a34f1ea
a31d70c3ebaa849175f09a20c50aa2eb8a2ce116 41f20cd8718f2d970ca8fbfbdb34f6e2b9ec89a9
4747c7cc0f9fb780ac173a81ae509a87b689a301 1288218f609b9434f1c177bd9c3b1b0d867b60da
4e63304fd7f0359fc6520023a96d8ef8f58e4455 7460c60ac89de3842808a3bc5e99ccde72898b37
)

[[ ${#blobs[@]} -eq 56 ]] || { echo "ERROR: expected 56 blob ids, got ${#blobs[@]}" >&2; exit 1; }
: > "$ARCHIVE"
for blob in "${blobs[@]}"; do
  [[ "$blob" =~ ^[0-9a-f]{40}$ ]] || { echo "ERROR: malformed blob id" >&2; exit 1; }
  gh api "repos/$GITHUB_REPOSITORY/git/blobs/$blob" --jq .content | tr -d '\n' | base64 --decode >> "$ARCHIVE"
done

echo "$EXPECTED_SHA256  $ARCHIVE" | sha256sum -c -
xz -t "$ARCHIVE"
tar -tJf "$ARCHIVE" > "$WORK_DIR/files.txt"
! grep -Eq '(^|/)\.\.(/|$)|^/' "$WORK_DIR/files.txt"
mkdir -p "$SOURCE"
tar -xJf "$ARCHIVE" -C "$SOURCE"
[[ "$(cat "$SOURCE/VERSION")" == "$EXPECTED_VERSION" ]]
for path in LICENSE README.md go.mod; do [[ -f "$SOURCE/$path" ]]; done
(
  cd "$SOURCE"
  go test ./... -count=1
  go vet ./...
  go test -race ./... -count=1
  go build -o /tmp/ltc ./cmd/ltc
  go build -o /tmp/ltc-ui ./cmd/ltc-ui
  bash -n packaging/linux/install-user.sh packaging/linux/uninstall-user.sh
)
echo "PASS: historical beta blob archive is complete, hash-verified, testable and buildable."
