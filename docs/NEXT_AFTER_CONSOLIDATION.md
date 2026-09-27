# Development restart order

After consolidation is complete, resume development from the canonical integration head in this order:

1. establish a complete buildable canonical application source tree;
2. run tests/static checks/build and packaging/install validation;
3. audit privacy/network/update behavior and dependency/supply-chain boundaries;
4. restore/verify desktop integration and easy Linux installation;
5. continue UX/customization development and effects work;
6. refresh public screenshots/metadata only from verified builds;
7. prepare a release candidate through a reviewed PR and explicit publication gate.

Do not restart feature work from an arbitrary historical branch.
