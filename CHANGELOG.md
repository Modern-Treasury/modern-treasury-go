# Changelog

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
