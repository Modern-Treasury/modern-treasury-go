# Changelog

## 2.0.0 (2023-10-19)

Full Changelog: [v1.5.2...v2.0.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v1.5.2...v2.0.0)

### ⚠ BREAKING CHANGES

* **types:** consolidate direction enums into a shared TransactionDirection type ([#107](https://github.com/Modern-Treasury/modern-treasury-go/issues/107))

### Features

* make webhook headers case insensitive ([#106](https://github.com/Modern-Treasury/modern-treasury-go/issues/106)) ([09d58d9](https://github.com/Modern-Treasury/modern-treasury-go/commit/09d58d9465d1f07b72e207a21c90b5652ad7847c))
* **types:** consolidate direction enums into a shared TransactionDirection type ([#107](https://github.com/Modern-Treasury/modern-treasury-go/issues/107)) ([2c32ee7](https://github.com/Modern-Treasury/modern-treasury-go/commit/2c32ee7829dd82cab93af1f830cf30d8f2ade1f6))


### Bug Fixes

* **api:** use date-time for effective_at ([#111](https://github.com/Modern-Treasury/modern-treasury-go/issues/111)) ([9987e10](https://github.com/Modern-Treasury/modern-treasury-go/commit/9987e10a5d30c293d187127e5e5a137cb36c01d3))


### Chores

* **internal:** rearrange client arguments ([#103](https://github.com/Modern-Treasury/modern-treasury-go/issues/103)) ([d93e977](https://github.com/Modern-Treasury/modern-treasury-go/commit/d93e977e6a8a9e5577031567ca9c75b97eb2232c))
* **internal:** reorder code ([#108](https://github.com/Modern-Treasury/modern-treasury-go/issues/108)) ([c6f25e4](https://github.com/Modern-Treasury/modern-treasury-go/commit/c6f25e4de17662ea119a5dd5a10c1da63a486b31))
* update README ([#102](https://github.com/Modern-Treasury/modern-treasury-go/issues/102)) ([5df84eb](https://github.com/Modern-Treasury/modern-treasury-go/commit/5df84eb61d34a5a37343065147a01ad60d031af0))


### Documentation

* **api.md:** improve formatting of webhook helpers ([#105](https://github.com/Modern-Treasury/modern-treasury-go/issues/105)) ([731f0c5](https://github.com/Modern-Treasury/modern-treasury-go/commit/731f0c564db9c6ae32fc33c2087d7f0092d87cb5))
* organisation -&gt; organization (UK to US English) ([#109](https://github.com/Modern-Treasury/modern-treasury-go/issues/109)) ([a15e655](https://github.com/Modern-Treasury/modern-treasury-go/commit/a15e6558a208a46d055603ee708b80a9939c3b02))

## 1.5.2 (2023-10-06)

Full Changelog: [v1.5.1...v1.5.2](https://github.com/modern-treasury/modern-treasury-go/compare/v1.5.1...v1.5.2)

### Bug Fixes

* prevent index out of range bug during auto-pagination ([#98](https://github.com/modern-treasury/modern-treasury-go/issues/98)) ([6db501f](https://github.com/modern-treasury/modern-treasury-go/commit/6db501f1ef403e1d80209723cec0289c94d9b832))

## 1.5.1 (2023-10-02)

Full Changelog: [v1.5.0...v1.5.1](https://github.com/Modern-Treasury/modern-treasury-go/compare/v1.5.0...v1.5.1)

## 1.5.0 (2023-09-25)

Full Changelog: [v1.4.1...v1.5.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v1.4.1...v1.5.0)

### Features

* **api:** updates ([#89](https://github.com/Modern-Treasury/modern-treasury-go/issues/89)) ([26d932e](https://github.com/Modern-Treasury/modern-treasury-go/commit/26d932e2178f3d823d673b14818504673ecccc0f))
* improve retry behavior on context deadline ([#92](https://github.com/Modern-Treasury/modern-treasury-go/issues/92)) ([852cca4](https://github.com/Modern-Treasury/modern-treasury-go/commit/852cca4c00794290dada40b85131e3ffa511dc72))
* retry on 408 Request Timeout ([#85](https://github.com/Modern-Treasury/modern-treasury-go/issues/85)) ([df0dce5](https://github.com/Modern-Treasury/modern-treasury-go/commit/df0dce5be8f34219e33fc10a4f314794a29faf17))


### Bug Fixes

* **client:** fix alias comment ([#90](https://github.com/Modern-Treasury/modern-treasury-go/issues/90)) ([8b55479](https://github.com/Modern-Treasury/modern-treasury-go/commit/8b554795e6e9b21b867c108887c70ac4282b2b11))
* **core:** improve retry behavior and related docs ([#86](https://github.com/Modern-Treasury/modern-treasury-go/issues/86)) ([06f6e59](https://github.com/Modern-Treasury/modern-treasury-go/commit/06f6e59e49138b494c08f7d346eb9a19fc02abe0))


### Chores

* **next => main:** release 1.4.1 ([240fd70](https://github.com/Modern-Treasury/modern-treasury-go/commit/240fd704fcfc1bfb6458b7023be02b76d6e7eab7))


### Documentation

* **api.md:** rename Top Level to client name ([#91](https://github.com/Modern-Treasury/modern-treasury-go/issues/91)) ([b2f1355](https://github.com/Modern-Treasury/modern-treasury-go/commit/b2f1355aae03ebac7d255c604ec016c58716e5c2))

## 1.4.1 (2023-09-12)

Full Changelog: [v1.4.0...v1.4.1](https://github.com/Modern-Treasury/modern-treasury-go/compare/v1.4.0...v1.4.1)

### Bug Fixes

* **core:** add null check to prevent segfault when canceling context ([#78](https://github.com/Modern-Treasury/modern-treasury-go/issues/78)) ([e120adf](https://github.com/Modern-Treasury/modern-treasury-go/commit/e120adf531e47e991cc6f51acd939eb89e4a2dc1))


### Chores

* **internal:** improve reliability of cancel delay test ([#80](https://github.com/Modern-Treasury/modern-treasury-go/issues/80)) ([53b87f8](https://github.com/Modern-Treasury/modern-treasury-go/commit/53b87f8535864878e36def5510150eae526f8493))
* **next => main:** release 1.4.1 ([f92bada](https://github.com/Modern-Treasury/modern-treasury-go/commit/f92bada0692b4574d545deead0885a3220a408fb))

## 1.4.1 (2023-09-12)

Full Changelog: [v1.4.0...v1.4.1](https://github.com/Modern-Treasury/modern-treasury-go/compare/v1.4.0...v1.4.1)

### Bug Fixes

* **core:** add null check to prevent segfault when canceling context ([#78](https://github.com/Modern-Treasury/modern-treasury-go/issues/78)) ([974d02b](https://github.com/Modern-Treasury/modern-treasury-go/commit/974d02b77b6cab8c646cac7b829836263eff264b))


### Chores

* **internal:** improve reliability of cancel delay test ([#80](https://github.com/Modern-Treasury/modern-treasury-go/issues/80)) ([0b74e83](https://github.com/Modern-Treasury/modern-treasury-go/commit/0b74e83b5ae1bf2b505b5de04158a0dc247bbee3))

## 1.4.0 (2023-09-06)

Full Changelog: [v1.3.0...v1.4.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v1.3.0...v1.4.0)

### Features

* fixes tests where an array has to have unique enum values ([#74](https://github.com/Modern-Treasury/modern-treasury-go/issues/74)) ([1b899fc](https://github.com/Modern-Treasury/modern-treasury-go/commit/1b899fc01e5f7e3ae1342f4fac2b63e42790e295))


### Documentation

* **readme:** add link to api.md ([#76](https://github.com/Modern-Treasury/modern-treasury-go/issues/76)) ([23bc6d7](https://github.com/Modern-Treasury/modern-treasury-go/commit/23bc6d7ea1a5625493ca618d6c91371f512cc127))

## 1.3.0 (2023-09-01)

Full Changelog: [v1.2.2...v1.3.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v1.2.2...v1.3.0)

### Features

* **types:** remove incorrectly included Idempotency-Key param ([#72](https://github.com/Modern-Treasury/modern-treasury-go/issues/72)) ([3f2fca4](https://github.com/Modern-Treasury/modern-treasury-go/commit/3f2fca4967dc623aeb87ac35e9db8dcd82d5a6d9))

## 1.2.2 (2023-08-31)

Full Changelog: [v1.2.1...v1.2.2](https://github.com/Modern-Treasury/modern-treasury-go/compare/v1.2.1...v1.2.2)

### Bug Fixes

* make paginated requests deserialize properly ([#70](https://github.com/Modern-Treasury/modern-treasury-go/issues/70)) ([0f59a32](https://github.com/Modern-Treasury/modern-treasury-go/commit/0f59a329645edb00998d93598fe5b1ab697699e0))

## 1.2.1 (2023-08-28)

Full Changelog: [v1.2.0...v1.2.1](https://github.com/Modern-Treasury/modern-treasury-go/compare/v1.2.0...v1.2.1)

### Chores

* assign default reviewers to release PRs ([#64](https://github.com/Modern-Treasury/modern-treasury-go/issues/64)) ([af8da66](https://github.com/Modern-Treasury/modern-treasury-go/commit/af8da6624d97e94e989cf825a836a1d1e6ac0e73))
* **ci:** setup workflows to create releases and release PRs ([#67](https://github.com/Modern-Treasury/modern-treasury-go/issues/67)) ([dcc4131](https://github.com/Modern-Treasury/modern-treasury-go/commit/dcc41318694ebe95a8770219517d24a834093bbc))

## [1.2.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v1.1.0...v1.2.0) (2023-08-11)


### Features

* **api:** add `metadata` in several places it was missing; add `description` ([#56](https://github.com/Modern-Treasury/modern-treasury-go/issues/56)) ([316baf5](https://github.com/Modern-Treasury/modern-treasury-go/commit/316baf52321911f81dfeb6cc73371818304f5ea7))
* **api:** support multiple `id`s in `ledger` `retrieve`/`list` endpoints ([#62](https://github.com/Modern-Treasury/modern-treasury-go/issues/62)) ([cb0455a](https://github.com/Modern-Treasury/modern-treasury-go/commit/cb0455a42195704d03c5df658171582688ca8664))
* **api:** updates ([#58](https://github.com/Modern-Treasury/modern-treasury-go/issues/58)) ([dc58726](https://github.com/Modern-Treasury/modern-treasury-go/commit/dc587260961a2641f8214699b54d1fa96fa3c568))


### Bug Fixes

* **client:** correctly set multipart form data boundary ([#60](https://github.com/Modern-Treasury/modern-treasury-go/issues/60)) ([871be0a](https://github.com/Modern-Treasury/modern-treasury-go/commit/871be0ac5216ea7fd1051736a8be4d06727429b9))


### Documentation

* **readme:** remove beta status + document versioning policy ([#59](https://github.com/Modern-Treasury/modern-treasury-go/issues/59)) ([7dcbabf](https://github.com/Modern-Treasury/modern-treasury-go/commit/7dcbabf9af5083d7d1767dbd46f2733d955ad4b4))

## [1.1.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v1.0.0...v2.0.0) (2023-08-01)


### ⚠ BREAKING CHANGES

* **types:** rename account connection flow to account collection flow ([#49](https://github.com/Modern-Treasury/modern-treasury-go/issues/49))
* **api:** update parameters for virtual account create request ([#48](https://github.com/Modern-Treasury/modern-treasury-go/issues/48))

### Features

* add `Bool` param field helper ([#51](https://github.com/Modern-Treasury/modern-treasury-go/issues/51)) ([a8d7b8f](https://github.com/Modern-Treasury/modern-treasury-go/commit/a8d7b8fbe64ea66a430d01b8a7ece472a35b3c5e))
* **api:** update parameters for virtual account create request ([#48](https://github.com/Modern-Treasury/modern-treasury-go/issues/48)) ([17ff81b](https://github.com/Modern-Treasury/modern-treasury-go/commit/17ff81b903acef27a9a08ee6d81f89a3d91df2ef))
* **api:** updates ([#53](https://github.com/Modern-Treasury/modern-treasury-go/issues/53)) ([a97abe0](https://github.com/Modern-Treasury/modern-treasury-go/commit/a97abe09b8157dcc2da60a8a31caa07d1b95fcf2))
* **api:** updates ([#54](https://github.com/Modern-Treasury/modern-treasury-go/issues/54)) ([9a45f0d](https://github.com/Modern-Treasury/modern-treasury-go/commit/9a45f0d502b798451fb544359a0c8ef8f086fde1))


### Refactors

* **types:** rename account connection flow to account collection flow ([#49](https://github.com/Modern-Treasury/modern-treasury-go/issues/49)) ([ca8dac5](https://github.com/Modern-Treasury/modern-treasury-go/commit/ca8dac5137e961d3e5a13e09712c9f82668e3fe8))


### Chores

* **internal:** add `codegen.log` to `.gitignore` ([#46](https://github.com/Modern-Treasury/modern-treasury-go/issues/46)) ([adab044](https://github.com/Modern-Treasury/modern-treasury-go/commit/adab044c40f27eb8aeac8ee1fa7cb906dc404217))
* **internal:** minor reformatting of code ([#55](https://github.com/Modern-Treasury/modern-treasury-go/issues/55)) ([48a3097](https://github.com/Modern-Treasury/modern-treasury-go/commit/48a30971393bbdf50da58b3c2064187376df026b))

## [1.0.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v0.0.1...v1.0.0) (2023-07-14)

### Chores

- unreleased changes ([#42](https://github.com/Modern-Treasury/modern-treasury-go/issues/42)) ([5685f72](https://github.com/Modern-Treasury/modern-treasury-go/commit/5685f723ca8dba82494dda6d876c814b49307ab8))
