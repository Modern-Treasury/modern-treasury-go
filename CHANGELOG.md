# Changelog

## 2.22.0 (2025-02-14)

Full Changelog: [v2.21.0...v2.22.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.21.0...v2.22.0)

### Features

* **client:** send `X-Stainless-Timeout` header ([#309](https://github.com/Modern-Treasury/modern-treasury-go/issues/309)) ([f17d1f0](https://github.com/Modern-Treasury/modern-treasury-go/commit/f17d1f0d635580e1c028e8b93ae4de9ffd536dd5))


### Bug Fixes

* **client:** don't truncate manually specified filenames ([#313](https://github.com/Modern-Treasury/modern-treasury-go/issues/313)) ([e6b36c0](https://github.com/Modern-Treasury/modern-treasury-go/commit/e6b36c0dfdbd547ab49d8c4118131bc9d746bfac))
* do not call path.Base on ContentType ([#312](https://github.com/Modern-Treasury/modern-treasury-go/issues/312)) ([4b12f61](https://github.com/Modern-Treasury/modern-treasury-go/commit/4b12f617665e5867992e4024261d7b0675f2b71e))
* fix early cancel when RequestTimeout is provided for streaming requests ([#311](https://github.com/Modern-Treasury/modern-treasury-go/issues/311)) ([3b6659a](https://github.com/Modern-Treasury/modern-treasury-go/commit/3b6659a7e7e1ce3faa6d44d68c404b17a0dafe4a))
* fix unicode encoding for json ([#306](https://github.com/Modern-Treasury/modern-treasury-go/issues/306)) ([c839501](https://github.com/Modern-Treasury/modern-treasury-go/commit/c83950183e03f6875b5d9dc44723e7518939116e))
* mark nullable property as nullable ([#305](https://github.com/Modern-Treasury/modern-treasury-go/issues/305)) ([7ca7290](https://github.com/Modern-Treasury/modern-treasury-go/commit/7ca7290d7403d4a98585b714e6d5b6264c811404))


### Chores

* add UnionUnmarshaler for responses that are interfaces ([#310](https://github.com/Modern-Treasury/modern-treasury-go/issues/310)) ([be8c5fc](https://github.com/Modern-Treasury/modern-treasury-go/commit/be8c5fc638f37dcbd36dfe2876f1faaddcdf5d0b))
* refactor client tests ([#303](https://github.com/Modern-Treasury/modern-treasury-go/issues/303)) ([1979afa](https://github.com/Modern-Treasury/modern-treasury-go/commit/1979afacd0aa200e7b8f0420fb7ace9e29a2a258))


### Documentation

* document raw responses ([#307](https://github.com/Modern-Treasury/modern-treasury-go/issues/307)) ([3fc66ac](https://github.com/Modern-Treasury/modern-treasury-go/commit/3fc66aca9b24f7cca87e8f850e80d249f165cb21))

## 2.21.0 (2025-01-22)

Full Changelog: [v2.20.5...v2.21.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.20.5...v2.21.0)

### Features

* support deprecated markers ([#301](https://github.com/Modern-Treasury/modern-treasury-go/issues/301)) ([7219c54](https://github.com/Modern-Treasury/modern-treasury-go/commit/7219c54536171a87ba9e01b5e14789529135b825))


### Bug Fixes

* fix apijson.Port for embedded structs ([#298](https://github.com/Modern-Treasury/modern-treasury-go/issues/298)) ([8e96a81](https://github.com/Modern-Treasury/modern-treasury-go/commit/8e96a81639555370d119762fface3884f81e674c))
* fix apijson.Port for embedded structs ([#300](https://github.com/Modern-Treasury/modern-treasury-go/issues/300)) ([9ffcd94](https://github.com/Modern-Treasury/modern-treasury-go/commit/9ffcd94df1229d7847b8efa5bb96af95eb03dc2c))


### Chores

* **api:** adds new APIs for LedgerAccountSettlement LedgerEntries ([#302](https://github.com/Modern-Treasury/modern-treasury-go/issues/302)) ([850744a](https://github.com/Modern-Treasury/modern-treasury-go/commit/850744aa75ee06026766dc843309d03dbdd11c46))

## 2.20.5 (2025-01-08)

Full Changelog: [v2.20.4...v2.20.5](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.20.4...v2.20.5)

### Chores

* bump license year ([#296](https://github.com/Modern-Treasury/modern-treasury-go/issues/296)) ([1def9db](https://github.com/Modern-Treasury/modern-treasury-go/commit/1def9db13f6be946ebb82a5a5fe760309d145a7a))


### Documentation

* **readme:** fix misplaced period ([#297](https://github.com/Modern-Treasury/modern-treasury-go/issues/297)) ([89b8913](https://github.com/Modern-Treasury/modern-treasury-go/commit/89b8913ba82887a5e8bcffbd27be1ca31e208c95))
* **readme:** fix typo ([#294](https://github.com/Modern-Treasury/modern-treasury-go/issues/294)) ([c7ac945](https://github.com/Modern-Treasury/modern-treasury-go/commit/c7ac94561a2177047acf5314b5e17b6bfaa24801))

## 2.20.4 (2024-11-25)

Full Changelog: [v2.20.3...v2.20.4](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.20.3...v2.20.4)

### Chores

* **api:** additional fields for requests to get BalanceReports and create LegalEntities ([#292](https://github.com/Modern-Treasury/modern-treasury-go/issues/292)) ([579f935](https://github.com/Modern-Treasury/modern-treasury-go/commit/579f935cba1451233b1112e266e0ce6f47b83e8e))
* sync openapi spec ([#293](https://github.com/Modern-Treasury/modern-treasury-go/issues/293)) ([7b6557d](https://github.com/Modern-Treasury/modern-treasury-go/commit/7b6557d35f239a3faf18f98d7ecadaef1e489615))
* **tests:** limit array example length ([#290](https://github.com/Modern-Treasury/modern-treasury-go/issues/290)) ([eb2208c](https://github.com/Modern-Treasury/modern-treasury-go/commit/eb2208cd5de0f0463da3e84f54f6325141096a9e))

## 2.20.3 (2024-11-11)

Full Changelog: [v2.20.2...v2.20.3](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.20.2...v2.20.3)

### Bug Fixes

* **api:** escape key values when encoding maps ([#287](https://github.com/Modern-Treasury/modern-treasury-go/issues/287)) ([c670a90](https://github.com/Modern-Treasury/modern-treasury-go/commit/c670a90f362d137a94eb68615e34fe3fdbe8bdb4))
* **client:** no panic on missing BaseURL ([#289](https://github.com/Modern-Treasury/modern-treasury-go/issues/289)) ([06234a3](https://github.com/Modern-Treasury/modern-treasury-go/commit/06234a320ec13620eafaab197526bcc079edbe73))
* correct required fields for flattened unions ([#288](https://github.com/Modern-Treasury/modern-treasury-go/issues/288)) ([7d582c5](https://github.com/Modern-Treasury/modern-treasury-go/commit/7d582c5252195e9c5b03259210c8ada9dd84ed06))


### Refactors

* sort fields for squashed union structs ([#285](https://github.com/Modern-Treasury/modern-treasury-go/issues/285)) ([79839fd](https://github.com/Modern-Treasury/modern-treasury-go/commit/79839fd7c1bbc28c0146e1c9217f682d46dbe0a2))

## 2.20.2 (2024-10-16)

Full Changelog: [v2.20.1...v2.20.2](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.20.1...v2.20.2)

### Features

* **api:** updates to required fields for ExpectedPayments ([#284](https://github.com/Modern-Treasury/modern-treasury-go/issues/284)) ([75e3ebc](https://github.com/Modern-Treasury/modern-treasury-go/commit/75e3ebc98bfb5d46e7875420b85e09e3521ca44e))
* move pagination package from internal to packages ([#283](https://github.com/Modern-Treasury/modern-treasury-go/issues/283)) ([5c19e96](https://github.com/Modern-Treasury/modern-treasury-go/commit/5c19e968872516c08c66367b97455e991365d69c))


### Chores

* fix GetNextPage docstring ([#281](https://github.com/Modern-Treasury/modern-treasury-go/issues/281)) ([acadfa3](https://github.com/Modern-Treasury/modern-treasury-go/commit/acadfa3987c5e31a7352b14b9f3b59f4016aa1ae))

## 2.20.1 (2024-10-09)

Full Changelog: [v2.20.0...v2.20.1](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.20.0...v2.20.1)

### Documentation

* improve and reference contributing documentation ([#279](https://github.com/Modern-Treasury/modern-treasury-go/issues/279)) ([b80d5c7](https://github.com/Modern-Treasury/modern-treasury-go/commit/b80d5c78651e090acbd5fa50cce7b0ee2e13371f))

## 2.20.0 (2024-09-24)

Full Changelog: [v2.19.0...v2.20.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.19.0...v2.20.0)

### Features

* **api:** add `usbank_payment_application_reference_id` to `reference_number_type` ([#275](https://github.com/Modern-Treasury/modern-treasury-go/issues/275)) ([dfa745b](https://github.com/Modern-Treasury/modern-treasury-go/commit/dfa745b1d54e56ea90ea63e6e9d0db24949b6431))
* **client:** send retry count header ([#278](https://github.com/Modern-Treasury/modern-treasury-go/issues/278)) ([299e838](https://github.com/Modern-Treasury/modern-treasury-go/commit/299e83800e095b62ca1f3bb017023668442acb62))


### Bug Fixes

* **requestconfig:** copy over more fields when cloning ([#273](https://github.com/Modern-Treasury/modern-treasury-go/issues/273)) ([39900c9](https://github.com/Modern-Treasury/modern-treasury-go/commit/39900c96bad9e2b9554f81dd830751a6e19728ff))


### Chores

* **api:** fields and parameters added to bulk actions, transactions and invoice creation ([#277](https://github.com/Modern-Treasury/modern-treasury-go/issues/277)) ([95ccfdf](https://github.com/Modern-Treasury/modern-treasury-go/commit/95ccfdf8fcfb0ef6f11ce44a014f0c466009677b))


### Documentation

* update CONTRIBUTING.md ([#276](https://github.com/Modern-Treasury/modern-treasury-go/issues/276)) ([e18b5e4](https://github.com/Modern-Treasury/modern-treasury-go/commit/e18b5e4417b66ec6f95d7099ce8b2566df1cbcaf))

## 2.19.0 (2024-09-09)

Full Changelog: [v2.18.0...v2.19.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.18.0...v2.19.0)

### Features

* **api:** add us_bank RTP ID's as reference_number_type ([#270](https://github.com/Modern-Treasury/modern-treasury-go/issues/270)) ([5508677](https://github.com/Modern-Treasury/modern-treasury-go/commit/5508677e3b84a3f424b0ff1df7347e5365611f60))


### Chores

* **docs:** update description of `bankgirot` to `se_bankgirot` ([#272](https://github.com/Modern-Treasury/modern-treasury-go/issues/272)) ([f11f19c](https://github.com/Modern-Treasury/modern-treasury-go/commit/f11f19c5a3d8d094dbd934e45161e906dabfd07e))

## 2.18.0 (2024-08-20)

Full Changelog: [v2.17.0...v2.18.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.17.0...v2.18.0)

### Features

* **api:** add wells fargo reference number types ([#268](https://github.com/Modern-Treasury/modern-treasury-go/issues/268)) ([56de966](https://github.com/Modern-Treasury/modern-treasury-go/commit/56de9661525389f49f15975d545077c69308be4a))

## 2.17.0 (2024-08-13)

Full Changelog: [v2.16.0...v2.17.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.16.0...v2.17.0)

### Features

* **api:** add pagination params 'created at' and 'updated at' ([#261](https://github.com/Modern-Treasury/modern-treasury-go/issues/261)) ([31177a6](https://github.com/Modern-Treasury/modern-treasury-go/commit/31177a6c449aa77d13d75936846db1a97038ee44))
* **api:** updates ([#267](https://github.com/Modern-Treasury/modern-treasury-go/issues/267)) ([62fac37](https://github.com/Modern-Treasury/modern-treasury-go/commit/62fac37bcc22e140999543ab6aa3d9ab06661c0a))


### Bug Fixes

* deserialization of struct unions that implement json.Unmarshaler ([#264](https://github.com/Modern-Treasury/modern-treasury-go/issues/264)) ([3b9fcd2](https://github.com/Modern-Treasury/modern-treasury-go/commit/3b9fcd230982fb949f7f396cccecd5a237a9620a))
* handle nil pagination responses when HTTP status is 200 ([#260](https://github.com/Modern-Treasury/modern-treasury-go/issues/260)) ([383d064](https://github.com/Modern-Treasury/modern-treasury-go/commit/383d0643c6c2699f0163abe9e39332ca6eafed8c))
* improve name for single-keyed union member ([#262](https://github.com/Modern-Treasury/modern-treasury-go/issues/262)) ([7df1ce3](https://github.com/Modern-Treasury/modern-treasury-go/commit/7df1ce39d034187754f053cf23b8e327c5270ae3))


### Chores

* bump Go to v1.21 ([#265](https://github.com/Modern-Treasury/modern-treasury-go/issues/265)) ([752d718](https://github.com/Modern-Treasury/modern-treasury-go/commit/752d7180072fcd237e6e19a2c89579574d078cbc))
* **ci:** bump prism mock server version ([#263](https://github.com/Modern-Treasury/modern-treasury-go/issues/263)) ([2202b23](https://github.com/Modern-Treasury/modern-treasury-go/commit/2202b23ef7dda975b55bd3caba553cab9d4410e3))
* **ci:** limit release doctor target branches ([#256](https://github.com/Modern-Treasury/modern-treasury-go/issues/256)) ([7e0563b](https://github.com/Modern-Treasury/modern-treasury-go/commit/7e0563b3118ebd2d96dd1fa8a7a80371b17e511b))
* **ci:** remove unused release doctor ([#258](https://github.com/Modern-Treasury/modern-treasury-go/issues/258)) ([33acece](https://github.com/Modern-Treasury/modern-treasury-go/commit/33aceceb2510f8d7ea56c152f210e4471fdd60ba))
* **examples:** minor formatting changes ([#266](https://github.com/Modern-Treasury/modern-treasury-go/issues/266)) ([144f4bd](https://github.com/Modern-Treasury/modern-treasury-go/commit/144f4bdb2445b3d311a88f6b2b7e5c50e346a5fc))
* **tests:** update prism version ([#259](https://github.com/Modern-Treasury/modern-treasury-go/issues/259)) ([7a03f6f](https://github.com/Modern-Treasury/modern-treasury-go/commit/7a03f6f54daa7a4c6106965122f82b2d32b5e012))

## 2.16.0 (2024-07-15)

Full Changelog: [v2.15.0...v2.16.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.15.0...v2.16.0)

### Features

* **api:** updates ([#255](https://github.com/Modern-Treasury/modern-treasury-go/issues/255)) ([7005f46](https://github.com/Modern-Treasury/modern-treasury-go/commit/7005f46e84906674c7c592a3018a209ca05ad53e))


### Bug Fixes

* **internal:** fix MarshalJSON logic for interface elemnets ([#251](https://github.com/Modern-Treasury/modern-treasury-go/issues/251)) ([4f27d96](https://github.com/Modern-Treasury/modern-treasury-go/commit/4f27d96a66c0f9876f87af6ead001a9bb163954c))
* use slice instead of appending to r.Options ([#249](https://github.com/Modern-Treasury/modern-treasury-go/issues/249)) ([8047c3a](https://github.com/Modern-Treasury/modern-treasury-go/commit/8047c3a7ff95cb464b8631f6fbcc6928c56a7835))


### Chores

* **ci:** also run workflows for PRs targeting `next` ([#252](https://github.com/Modern-Treasury/modern-treasury-go/issues/252)) ([e213167](https://github.com/Modern-Treasury/modern-treasury-go/commit/e213167de728431b427f9f7794b30dac257fd71d))
* **internal:** improve deserialization of embedded structs ([#250](https://github.com/Modern-Treasury/modern-treasury-go/issues/250)) ([907dd05](https://github.com/Modern-Treasury/modern-treasury-go/commit/907dd0578d6cb18bb74a522894b5c112bde39ae8))
* **internal:** minor changes to tests ([#254](https://github.com/Modern-Treasury/modern-treasury-go/issues/254)) ([4f8c716](https://github.com/Modern-Treasury/modern-treasury-go/commit/4f8c716521b4209e03c8cc0a4057fe5a1568da51))


### Documentation

* add better documentation for flattened enum types ([#246](https://github.com/Modern-Treasury/modern-treasury-go/issues/246)) ([e8612db](https://github.com/Modern-Treasury/modern-treasury-go/commit/e8612db6d59df6255b620de9a26bc19c7e087752))
* **examples:** update example values ([#253](https://github.com/Modern-Treasury/modern-treasury-go/issues/253)) ([3af4284](https://github.com/Modern-Treasury/modern-treasury-go/commit/3af4284341c4b147b86b55cb6c4c14273a7d8d7b))


### Refactors

* rename of inconsistent union variants ([#248](https://github.com/Modern-Treasury/modern-treasury-go/issues/248)) ([7851b60](https://github.com/Modern-Treasury/modern-treasury-go/commit/7851b60cf9b8a27e0e3737ffc6d60af8088218e2))

## 2.15.0 (2024-07-01)

Full Changelog: [v2.14.0...v2.15.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.14.0...v2.15.0)

### Features

* **api:** updates ([#245](https://github.com/Modern-Treasury/modern-treasury-go/issues/245)) ([877f9c3](https://github.com/Modern-Treasury/modern-treasury-go/commit/877f9c33a4a7a6bb9675f4093a23b8d2ead40df7))


### Bug Fixes

* fix ExtraFields serialization / deserialization ([#243](https://github.com/Modern-Treasury/modern-treasury-go/issues/243)) ([8b82674](https://github.com/Modern-Treasury/modern-treasury-go/commit/8b8267465a9cc6dec123e1c4c9ef17e6d82d22bf))
* fix port function for interface{} types ([#241](https://github.com/Modern-Treasury/modern-treasury-go/issues/241)) ([17ad20a](https://github.com/Modern-Treasury/modern-treasury-go/commit/17ad20a6216f61282a847aba55ea09acd20b33fd))


### Chores

* gitignore test server logs ([#244](https://github.com/Modern-Treasury/modern-treasury-go/issues/244)) ([268f263](https://github.com/Modern-Treasury/modern-treasury-go/commit/268f263ae34526753e8800009d0e2446dcf4b868))

## 2.14.0 (2024-06-05)

Full Changelog: [v2.13.1...v2.14.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.13.1...v2.14.0)

### Features

* **api:** add kr_brn kr_crn kr_rrn enum values ([#235](https://github.com/Modern-Treasury/modern-treasury-go/issues/235)) ([5125730](https://github.com/Modern-Treasury/modern-treasury-go/commit/512573085593da6b7e5072bfc8c5be832940ea03))
* **api:** add risk rating field ([#237](https://github.com/Modern-Treasury/modern-treasury-go/issues/237)) ([3c4e08b](https://github.com/Modern-Treasury/modern-treasury-go/commit/3c4e08b626fb178182109d41d6ebfeabd7969a22))


### Bug Fixes

* fix enum type to be non nullable ([#240](https://github.com/Modern-Treasury/modern-treasury-go/issues/240)) ([b7c4130](https://github.com/Modern-Treasury/modern-treasury-go/commit/b7c413090095a47b330a3e798c61bf281109ba47))
* **internal:** fix the way that unions are deserialized in nested arrays ([#238](https://github.com/Modern-Treasury/modern-treasury-go/issues/238)) ([abf98c0](https://github.com/Modern-Treasury/modern-treasury-go/commit/abf98c0192f82d0c68cc300838b45f5638b4bf19))


### Chores

* **internal:** sync urls ([#239](https://github.com/Modern-Treasury/modern-treasury-go/issues/239)) ([535fd43](https://github.com/Modern-Treasury/modern-treasury-go/commit/535fd432cc645f1dab1266647c02ce7cdac3f31c))

## 2.13.1 (2024-05-30)

Full Changelog: [v2.13.0...v2.13.1](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.13.0...v2.13.1)

### Chores

* **internal:** add code reviewer ([#233](https://github.com/Modern-Treasury/modern-treasury-go/issues/233)) ([9bc0ebc](https://github.com/Modern-Treasury/modern-treasury-go/commit/9bc0ebceaf7a1dc3b5f500f8c280a05d693e363e))

## 2.13.0 (2024-05-28)

Full Changelog: [v2.12.0...v2.13.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.12.0...v2.13.0)

### Features

* **api:** add currency to ledger account categories ([#228](https://github.com/Modern-Treasury/modern-treasury-go/issues/228)) ([76766e6](https://github.com/Modern-Treasury/modern-treasury-go/commit/76766e64f19092f9e3eb59e340e434127c4a7c05))
* **api:** invoice overdue reminders ([f9d2168](https://github.com/Modern-Treasury/modern-treasury-go/commit/f9d21682aadcaaa082d38f9556b50fa807a5e13a))
* **api:** mark ConnectionLegalEntity response properties as required ([#231](https://github.com/Modern-Treasury/modern-treasury-go/issues/231)) ([ce49af3](https://github.com/Modern-Treasury/modern-treasury-go/commit/ce49af3838e8a154af6b31bba55e23cfdc7ab384))
* **api:** remove deprecated ledger account payouts ([#229](https://github.com/Modern-Treasury/modern-treasury-go/issues/229)) ([f9d2168](https://github.com/Modern-Treasury/modern-treasury-go/commit/f9d21682aadcaaa082d38f9556b50fa807a5e13a))
* **api:** updates ([#219](https://github.com/Modern-Treasury/modern-treasury-go/issues/219)) ([a3e27cd](https://github.com/Modern-Treasury/modern-treasury-go/commit/a3e27cd22b2e944540853568fe1583716c415086))
* **api:** various API updates ([#215](https://github.com/Modern-Treasury/modern-treasury-go/issues/215)) ([5ee8bc2](https://github.com/Modern-Treasury/modern-treasury-go/commit/5ee8bc2adb9c238d89dc11a0708a2531fb3fd401))
* better validation of path params ([#230](https://github.com/Modern-Treasury/modern-treasury-go/issues/230)) ([a4d987b](https://github.com/Modern-Treasury/modern-treasury-go/commit/a4d987b389bc8f37d6a879f876afcec1773b11c9))
* propagate resource description field from stainless config to SDK docs ([#222](https://github.com/Modern-Treasury/modern-treasury-go/issues/222)) ([511b951](https://github.com/Modern-Treasury/modern-treasury-go/commit/511b9513882f1f236fecbb8aea75fa055dfd06d3))


### Bug Fixes

* fix reading the error body more than once ([#226](https://github.com/Modern-Treasury/modern-treasury-go/issues/226)) ([b5294a2](https://github.com/Modern-Treasury/modern-treasury-go/commit/b5294a2113e61cae5f0e512b214c2f1097e606e7))
* make shared package public ([#214](https://github.com/Modern-Treasury/modern-treasury-go/issues/214)) ([a6cc5de](https://github.com/Modern-Treasury/modern-treasury-go/commit/a6cc5de7e8bfd73c322bdd19f72dcd39edf43574))
* **test:** fix test github actions job ([#218](https://github.com/Modern-Treasury/modern-treasury-go/issues/218)) ([3589231](https://github.com/Modern-Treasury/modern-treasury-go/commit/358923185f574c9c4221493a611ef22e8deca706))


### Chores

* **docs:** add SECURITY.md ([#223](https://github.com/Modern-Treasury/modern-treasury-go/issues/223)) ([5813f5d](https://github.com/Modern-Treasury/modern-treasury-go/commit/5813f5dfb2febe924b0b369b2ef942340fc679c0))
* **docs:** streamline payment purpose and vendor failure handling ([#224](https://github.com/Modern-Treasury/modern-treasury-go/issues/224)) ([4e1b6b9](https://github.com/Modern-Treasury/modern-treasury-go/commit/4e1b6b9fb84713fd2968a064cdb9a64afbba48eb))
* **internal:** add link to openapi spec ([#216](https://github.com/Modern-Treasury/modern-treasury-go/issues/216)) ([d0fa4a7](https://github.com/Modern-Treasury/modern-treasury-go/commit/d0fa4a7f25fb068af38ecbdbbaa1893cd761ed6f))
* **internal:** add scripts/test, scripts/mock and add ci job ([#217](https://github.com/Modern-Treasury/modern-treasury-go/issues/217)) ([1baf31a](https://github.com/Modern-Treasury/modern-treasury-go/commit/1baf31aa1ad5eb180f0caf872e3162495a74cc1d))
* **internal:** add slightly better logging to scripts ([#225](https://github.com/Modern-Treasury/modern-treasury-go/issues/225)) ([2133049](https://github.com/Modern-Treasury/modern-treasury-go/commit/2133049a3049af7d1700216f13ef802ecaa2c32d))
* **internal:** fix bootstrap script ([#220](https://github.com/Modern-Treasury/modern-treasury-go/issues/220)) ([2f72aa6](https://github.com/Modern-Treasury/modern-treasury-go/commit/2f72aa6cc5d485af47330775db9dd71a1ce75542))
* **internal:** fix format script ([#232](https://github.com/Modern-Treasury/modern-treasury-go/issues/232)) ([667e918](https://github.com/Modern-Treasury/modern-treasury-go/commit/667e9184f8369b5733eef660552a6d813169ff07))
* **internal:** fix Port function for number and boolean enums ([#213](https://github.com/Modern-Treasury/modern-treasury-go/issues/213)) ([1e8e1dd](https://github.com/Modern-Treasury/modern-treasury-go/commit/1e8e1ddb59515144445d6356c7723bfdd1caa91b))
* **internal:** support parsing other json content types ([#227](https://github.com/Modern-Treasury/modern-treasury-go/issues/227)) ([09e137d](https://github.com/Modern-Treasury/modern-treasury-go/commit/09e137d010065c7547a0cc92978fd1e856d3dc85))
* **internal:** use actions/checkout@v4 for codeflow ([#211](https://github.com/Modern-Treasury/modern-treasury-go/issues/211)) ([22627d1](https://github.com/Modern-Treasury/modern-treasury-go/commit/22627d15a8b4d878524efb2c99e7bc05574fb8db))

## 2.12.0 (2024-04-17)

Full Changelog: [v2.11.0...v2.12.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.11.0...v2.12.0)

### Features

* **api:** add ledger_transaction_id field to reversal ([#210](https://github.com/Modern-Treasury/modern-treasury-go/issues/210)) ([8a812e4](https://github.com/Modern-Treasury/modern-treasury-go/commit/8a812e4483c596f79ca446672b70e49e171cd235))
* **option:** add option to provide a raw request body ([#208](https://github.com/Modern-Treasury/modern-treasury-go/issues/208)) ([d4f6079](https://github.com/Modern-Treasury/modern-treasury-go/commit/d4f6079694ec83125f78b389429c58578812ac44))


### Chores

* change test names ([#206](https://github.com/Modern-Treasury/modern-treasury-go/issues/206)) ([edfd0ca](https://github.com/Modern-Treasury/modern-treasury-go/commit/edfd0ca9d4b3d7cd08a928be9008ae2d3b0ede1d))
* **internal:** formatting ([#207](https://github.com/Modern-Treasury/modern-treasury-go/issues/207)) ([d79bf8f](https://github.com/Modern-Treasury/modern-treasury-go/commit/d79bf8fa2300eceb922b186dc75f5aee2a4d5c31))


### Documentation

* **examples:** use counterparties in snippets ([#204](https://github.com/Modern-Treasury/modern-treasury-go/issues/204)) ([fef99d7](https://github.com/Modern-Treasury/modern-treasury-go/commit/fef99d7b1a771f15c30aaba8962a171749320b76))


### Build System

* configure UTF-8 locale in devcontainer ([#209](https://github.com/Modern-Treasury/modern-treasury-go/issues/209)) ([6bb7390](https://github.com/Modern-Treasury/modern-treasury-go/commit/6bb73903cdc9410df32e81bf19460045bfd06ed4))

## 2.11.0 (2024-04-04)

Full Changelog: [v2.10.0...v2.11.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.10.0...v2.11.0)

### Features

* **api:** add id type in_lei ([#202](https://github.com/Modern-Treasury/modern-treasury-go/issues/202)) ([10c797a](https://github.com/Modern-Treasury/modern-treasury-go/commit/10c797aa3794979f635344d802c972a6bfc9f955))
* **api:** update account number type enum ([#199](https://github.com/Modern-Treasury/modern-treasury-go/issues/199)) ([dac9771](https://github.com/Modern-Treasury/modern-treasury-go/commit/dac9771b237da84a99853e0c1fdced9cf42c4de6))
* **client:** implement raw requests methods on client ([#196](https://github.com/Modern-Treasury/modern-treasury-go/issues/196)) ([bd07bec](https://github.com/Modern-Treasury/modern-treasury-go/commit/bd07bec830587bd6daa7944a45040cb3c861d489))


### Bug Fixes

* adjust how bulk request items are accessed ([#203](https://github.com/Modern-Treasury/modern-treasury-go/issues/203)) ([4d87a40](https://github.com/Modern-Treasury/modern-treasury-go/commit/4d87a4055f5e9e17acb674434d4a4256581ebc19))


### Chores

* **internal:** move pagination types to pagination package ([#198](https://github.com/Modern-Treasury/modern-treasury-go/issues/198)) ([9cdf8a5](https://github.com/Modern-Treasury/modern-treasury-go/commit/9cdf8a5d213ff42690943e5cae4ba2ad4a200e9d))
* **internal:** use a time zone less likely to conflict with the local one ([#201](https://github.com/Modern-Treasury/modern-treasury-go/issues/201)) ([b8dc880](https://github.com/Modern-Treasury/modern-treasury-go/commit/b8dc8807642a0f5e986dc834a136b764b5a4667b))

## 2.10.0 (2024-03-26)

Full Changelog: [v2.9.0...v2.10.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.9.0...v2.10.0)

### Features

* add IsKnown method to enums ([#187](https://github.com/Modern-Treasury/modern-treasury-go/issues/187)) ([aab3062](https://github.com/Modern-Treasury/modern-treasury-go/commit/aab30629851e95d32a1e46f97320a008615cb415))
* **api:** add citibank enums ([#192](https://github.com/Modern-Treasury/modern-treasury-go/issues/192)) ([a20e06f](https://github.com/Modern-Treasury/modern-treasury-go/commit/a20e06f3a9a559fc17141c14a3b630cea6a80a3a))
* **api:** add date_formed property to legal entities ([#181](https://github.com/Modern-Treasury/modern-treasury-go/issues/181)) ([f91bb6d](https://github.com/Modern-Treasury/modern-treasury-go/commit/f91bb6df8624cec1747c4bd16d988fc894f08128))
* **api:** add line item metadata ([cb9243b](https://github.com/Modern-Treasury/modern-treasury-go/commit/cb9243b590546354d1bcf2c07c1b68c153a24094))
* **api:** extend list invoices query params ([#184](https://github.com/Modern-Treasury/modern-treasury-go/issues/184)) ([2512a6e](https://github.com/Modern-Treasury/modern-treasury-go/commit/2512a6e1e870ec58b6e89a24c3a21f2ebbcf7e56))
* **api:** introduce bulk transaction create ([#190](https://github.com/Modern-Treasury/modern-treasury-go/issues/190)) ([cb9243b](https://github.com/Modern-Treasury/modern-treasury-go/commit/cb9243b590546354d1bcf2c07c1b68c153a24094))
* **api:** rename `associated_legal_entity` to `child_legal_entity` ([#188](https://github.com/Modern-Treasury/modern-treasury-go/issues/188)) ([0ec3815](https://github.com/Modern-Treasury/modern-treasury-go/commit/0ec3815f5aec94e1b29745343fe32c73c27b7b54))
* **api:** rename `id_type` enum from `cl_nut` to `cl_rut` ([0ec3815](https://github.com/Modern-Treasury/modern-treasury-go/commit/0ec3815f5aec94e1b29745343fe32c73c27b7b54))
* **api:** updates ([#191](https://github.com/Modern-Treasury/modern-treasury-go/issues/191)) ([7a6d391](https://github.com/Modern-Treasury/modern-treasury-go/commit/7a6d391ffacb06814822aee7a8604a4c8c6ab028))
* set user-agent header by default when making requests ([#183](https://github.com/Modern-Treasury/modern-treasury-go/issues/183)) ([f6a57ad](https://github.com/Modern-Treasury/modern-treasury-go/commit/f6a57ad21dc928995c6906e23c78c54676cf0fe0))


### Chores

* add back removed code ([6222d76](https://github.com/Modern-Treasury/modern-treasury-go/commit/6222d76024157768145265734dbca0838e56886a))
* **internal:** temporary commit ([#194](https://github.com/Modern-Treasury/modern-treasury-go/issues/194)) ([0d3ba33](https://github.com/Modern-Treasury/modern-treasury-go/commit/0d3ba334faf79bb45928f9554e483a992a43f75f))
* **internal:** update generated pragma comment ([#186](https://github.com/Modern-Treasury/modern-treasury-go/issues/186)) ([3745ce4](https://github.com/Modern-Treasury/modern-treasury-go/commit/3745ce4fa36138e29ba15cac9cc1b6a79dd87a6b))


### Documentation

* fix typo in CONTRIBUTING.md ([#185](https://github.com/Modern-Treasury/modern-treasury-go/issues/185)) ([8de2177](https://github.com/Modern-Treasury/modern-treasury-go/commit/8de2177a9a5928cd7b8a5c254f85707ef360de83))
* fix typo in docstring for Null() ([#195](https://github.com/Modern-Treasury/modern-treasury-go/issues/195)) ([b3fdb63](https://github.com/Modern-Treasury/modern-treasury-go/commit/b3fdb63095c79c4da7ed5a445a6bfef02faf122a))
* **readme:** consistent use of sentence case in headings ([#189](https://github.com/Modern-Treasury/modern-treasury-go/issues/189)) ([a7df734](https://github.com/Modern-Treasury/modern-treasury-go/commit/a7df73427293ffae6ebee889ba755419bb00fa1f))
* **readme:** document file uploads ([#193](https://github.com/Modern-Treasury/modern-treasury-go/issues/193)) ([a6b3c71](https://github.com/Modern-Treasury/modern-treasury-go/commit/a6b3c717be407cd85d6a082a2540be59b846b1ab))

## 2.9.0 (2024-03-12)

Full Changelog: [v2.8.0...v2.9.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.8.0...v2.9.0)

### Features

* **api:** add legal_structure enum member ([#173](https://github.com/Modern-Treasury/modern-treasury-go/issues/173)) ([a587b8b](https://github.com/Modern-Treasury/modern-treasury-go/commit/a587b8bbe87231ddadd665fb6766145c56b7ea7f))
* implement public RawJSON of response structs ([#175](https://github.com/Modern-Treasury/modern-treasury-go/issues/175)) ([856c7db](https://github.com/Modern-Treasury/modern-treasury-go/commit/856c7dbd40a0091848b74ba762dcbffaba1d8667))


### Bug Fixes

* **client:** don't include ? in path unless necessary ([#180](https://github.com/Modern-Treasury/modern-treasury-go/issues/180)) ([d07c304](https://github.com/Modern-Treasury/modern-treasury-go/commit/d07c30429a86cdb397666a9832e13999b12da344))
* fix String() behavior of param.Field ([#179](https://github.com/Modern-Treasury/modern-treasury-go/issues/179)) ([9fd583a](https://github.com/Modern-Treasury/modern-treasury-go/commit/9fd583a91d7ef10588edee405bd0d5d86a339c38))
* fix union deserialization for multiple objects ([#176](https://github.com/Modern-Treasury/modern-treasury-go/issues/176)) ([8e03bd0](https://github.com/Modern-Treasury/modern-treasury-go/commit/8e03bd0e7fbfc78b0878b578c1081ed873a2ac02))


### Chores

* **internal:** improve union deserialization logic ([#177](https://github.com/Modern-Treasury/modern-treasury-go/issues/177)) ([29b1769](https://github.com/Modern-Treasury/modern-treasury-go/commit/29b17693e1d4a31aba478c92157c4283227ea2d6))


### Documentation

* **contributing:** add a CONTRIBUTING.md ([#178](https://github.com/Modern-Treasury/modern-treasury-go/issues/178)) ([714b7f9](https://github.com/Modern-Treasury/modern-treasury-go/commit/714b7f9330e5aaaf72e8e456307e59384949ea9b))

## 2.8.0 (2024-02-29)

Full Changelog: [v2.7.0...v2.8.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.7.0...v2.8.0)

### Features

* **api:** add legal_entities resource ([#172](https://github.com/Modern-Treasury/modern-treasury-go/issues/172)) ([8fa4207](https://github.com/Modern-Treasury/modern-treasury-go/commit/8fa42079c41b2367d67018b00449ba5429151af2))
* **api:** added foreign exchange rate information ([#169](https://github.com/Modern-Treasury/modern-treasury-go/issues/169)) ([3d10c72](https://github.com/Modern-Treasury/modern-treasury-go/commit/3d10c72fdd0f3f6e053bdac97ba46eb95acced00))


### Chores

* **ci:** uses Stainless GitHub App for releases ([#165](https://github.com/Modern-Treasury/modern-treasury-go/issues/165)) ([56829ab](https://github.com/Modern-Treasury/modern-treasury-go/commit/56829ab4fe666fd647fc8afe58b6dc6919200c08))
* **internal:** refactor release environment script ([#167](https://github.com/Modern-Treasury/modern-treasury-go/issues/167)) ([04be1be](https://github.com/Modern-Treasury/modern-treasury-go/commit/04be1bed1be805a5f6c4602c016f1ee50bced38b))
* **internal:** update deps ([#170](https://github.com/Modern-Treasury/modern-treasury-go/issues/170)) ([b900168](https://github.com/Modern-Treasury/modern-treasury-go/commit/b9001683dbed621456e66db4caec15425b52e82f))


### Documentation

* **readme:** improve wording ([#171](https://github.com/Modern-Treasury/modern-treasury-go/issues/171)) ([74944f4](https://github.com/Modern-Treasury/modern-treasury-go/commit/74944f48ecd77a36945f6ce63144cbd3adbf55c5))

## 2.7.0 (2024-02-13)

Full Changelog: [v2.6.1...v2.7.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.6.1...v2.7.0)

### Features

* **api:** updates ([#164](https://github.com/Modern-Treasury/modern-treasury-go/issues/164)) ([4e1de14](https://github.com/Modern-Treasury/modern-treasury-go/commit/4e1de1471f22a08682aa2920fca6a39038372d56))
* **api:** updates parameters, error codes ([#160](https://github.com/Modern-Treasury/modern-treasury-go/issues/160)) ([bc91638](https://github.com/Modern-Treasury/modern-treasury-go/commit/bc9163872ad62ed94d16f78dc8eef64c9369a2ae))


### Chores

* **interal:** make link to api.md relative ([#161](https://github.com/Modern-Treasury/modern-treasury-go/issues/161)) ([7777eeb](https://github.com/Modern-Treasury/modern-treasury-go/commit/7777eeb96bcc36d9ed7f0ed12a9b6dd5348fc391))
* **internal:** adjust formatting ([#162](https://github.com/Modern-Treasury/modern-treasury-go/issues/162)) ([4162967](https://github.com/Modern-Treasury/modern-treasury-go/commit/4162967fd03ebaa3f88483516ac4af3a76acec49))
* **internal:** bump timeout threshold in tests ([#163](https://github.com/Modern-Treasury/modern-treasury-go/issues/163)) ([a78a751](https://github.com/Modern-Treasury/modern-treasury-go/commit/a78a751469ef3200c20f0f9088b26b6f43e6c189))
* **internal:** support pre-release versioning ([#158](https://github.com/Modern-Treasury/modern-treasury-go/issues/158)) ([89392d2](https://github.com/Modern-Treasury/modern-treasury-go/commit/89392d2a515a8b5060c1358603da76f0288971f9))

## 2.6.1 (2024-01-30)

Full Changelog: [v2.6.0...v2.6.1](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.6.0...v2.6.1)

### Bug Fixes

* **ci:** ignore stainless-app edits to release PR title ([#156](https://github.com/Modern-Treasury/modern-treasury-go/issues/156)) ([4415794](https://github.com/Modern-Treasury/modern-treasury-go/commit/44157943346e3a7f1ed02ad90c0bdb4352bbeebc))
* **test:** avoid test failures when SKIP_MOCK_TESTS is not set ([#155](https://github.com/Modern-Treasury/modern-treasury-go/issues/155)) ([10c1db0](https://github.com/Modern-Treasury/modern-treasury-go/commit/10c1db047adc3919d4e85bedef65ae32ea06b9c7))


### Chores

* **internal:** parse date-time strings more leniently ([#157](https://github.com/Modern-Treasury/modern-treasury-go/issues/157)) ([39a761b](https://github.com/Modern-Treasury/modern-treasury-go/commit/39a761b427b9795890a56902e49feda631104140))
* **internal:** speculative retry-after-ms support ([#153](https://github.com/Modern-Treasury/modern-treasury-go/issues/153)) ([324e894](https://github.com/Modern-Treasury/modern-treasury-go/commit/324e8949183ac91987e3249d9044257f56299514))

## 2.6.0 (2024-01-16)

Full Changelog: [v2.5.0...v2.6.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.5.0...v2.6.0)

### Features

* **api:** add `ledger_transactions` to expected payment request ([#149](https://github.com/Modern-Treasury/modern-treasury-go/issues/149)) ([cab1d97](https://github.com/Modern-Treasury/modern-treasury-go/commit/cab1d97ab70209985001a3d64719e7799e2ac150))
* **api:** add create and delete operations for internal accounts balance reports ([#151](https://github.com/Modern-Treasury/modern-treasury-go/issues/151)) ([6e54baf](https://github.com/Modern-Treasury/modern-treasury-go/commit/6e54baf76970a577ec5be9fd999067260af80150))


### Chores

* add .keep files for examples and custom code directories ([#150](https://github.com/Modern-Treasury/modern-treasury-go/issues/150)) ([8a18c2c](https://github.com/Modern-Treasury/modern-treasury-go/commit/8a18c2cb197edc914c3f4441bbc9e691b14a1b21))


### Documentation

* **readme:** improve api reference ([#152](https://github.com/Modern-Treasury/modern-treasury-go/issues/152)) ([6bdb2b8](https://github.com/Modern-Treasury/modern-treasury-go/commit/6bdb2b82efd7b97e5e1782ec31ec2c28da7f60f2))


### Refactors

* remove excess whitespace ([#147](https://github.com/Modern-Treasury/modern-treasury-go/issues/147)) ([8f49166](https://github.com/Modern-Treasury/modern-treasury-go/commit/8f491663b473304bf57b85f0fbaa7bb2153ba7e6))

## 2.5.0 (2024-01-02)

Full Changelog: [v2.4.1...v2.5.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.4.1...v2.5.0)

### Features

* **api:** remove reversed and reversing ledger account status type ([#144](https://github.com/Modern-Treasury/modern-treasury-go/issues/144)) ([c199f39](https://github.com/Modern-Treasury/modern-treasury-go/commit/c199f39cc7d6008bffeb43a6d6797b0dfac6d21b))
* **api:** updates ([#141](https://github.com/Modern-Treasury/modern-treasury-go/issues/141)) ([3de6b78](https://github.com/Modern-Treasury/modern-treasury-go/commit/3de6b787f8886c3e020328c9ee57bafb4aa38df9))
* **internal:** fallback to json serialization if no serialization methods are defined ([#142](https://github.com/Modern-Treasury/modern-treasury-go/issues/142)) ([650bd85](https://github.com/Modern-Treasury/modern-treasury-go/commit/650bd854e309c487862a0c39e73389b53d6b7b4c))


### Chores

* **ci:** run release workflow once per day ([#143](https://github.com/Modern-Treasury/modern-treasury-go/issues/143)) ([3a93291](https://github.com/Modern-Treasury/modern-treasury-go/commit/3a93291b2a6f2704357cdfc6c0360f41da614444))
* **internal:** bump license ([#146](https://github.com/Modern-Treasury/modern-treasury-go/issues/146)) ([9f0d312](https://github.com/Modern-Treasury/modern-treasury-go/commit/9f0d312dbe06baa875e6fede7cefd0cac89e5c47))


### Documentation

* **options:** fix link to readme ([#145](https://github.com/Modern-Treasury/modern-treasury-go/issues/145)) ([6bb5888](https://github.com/Modern-Treasury/modern-treasury-go/commit/6bb588809ce63e31d97374388acebd1a0cfbd27e))

## 2.4.1 (2023-12-04)

Full Changelog: [v2.4.0...v2.4.1](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.4.0...v2.4.1)

## 2.4.0 (2023-11-21)

Full Changelog: [v2.3.0...v2.4.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.3.0...v2.4.0)

### Features

* **api:** updates ([#136](https://github.com/Modern-Treasury/modern-treasury-go/issues/136)) ([c8e98e0](https://github.com/Modern-Treasury/modern-treasury-go/commit/c8e98e087d8594399014cab00cc9974e4cfbe0fd))


### Bug Fixes

* stop sending default idempotency headers with GET requests ([#134](https://github.com/Modern-Treasury/modern-treasury-go/issues/134)) ([2372568](https://github.com/Modern-Treasury/modern-treasury-go/commit/2372568abfa9fd5f71cb7a28b9cd034442d4b407))

## 2.3.0 (2023-11-16)

Full Changelog: [v2.2.0...v2.3.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.2.0...v2.3.0)

### Features

* **api:** add dk_interbank_clearing_code and dk_nets enum members ([#131](https://github.com/Modern-Treasury/modern-treasury-go/issues/131)) ([913f398](https://github.com/Modern-Treasury/modern-treasury-go/commit/913f3987ada5295def5f161bbb3907dc26ecf481))
* **api:** updates ([#128](https://github.com/Modern-Treasury/modern-treasury-go/issues/128)) ([8210a29](https://github.com/Modern-Treasury/modern-treasury-go/commit/8210a29fb80e73f252559bd12c90ddf699f7f6f3))


### Bug Fixes

* make options.WithHeader utils case-insensitive ([#130](https://github.com/Modern-Treasury/modern-treasury-go/issues/130)) ([52b942e](https://github.com/Modern-Treasury/modern-treasury-go/commit/52b942e2579754fc03629e80b3dbfa90818ea880))


### Chores

* **internal:** skip bulk requests tests ([#132](https://github.com/Modern-Treasury/modern-treasury-go/issues/132)) ([e5df103](https://github.com/Modern-Treasury/modern-treasury-go/commit/e5df1036b95b49200fdacbb90aa310432c1a65a0))


### Refactors

* do not include `JSON` fields when serialising back to json ([#133](https://github.com/Modern-Treasury/modern-treasury-go/issues/133)) ([b54fb6c](https://github.com/Modern-Treasury/modern-treasury-go/commit/b54fb6cbfb74000c512c71ffb5d0a15d8a765244))

## 2.2.0 (2023-11-05)

Full Changelog: [v2.1.0...v2.2.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.1.0...v2.2.0)

### Features

* **client:** allow binary returns ([#125](https://github.com/Modern-Treasury/modern-treasury-go/issues/125)) ([68970b0](https://github.com/Modern-Treasury/modern-treasury-go/commit/68970b06367d463875503e04bcfb6d3ed811bd8b))


### Documentation

* **readme:** improve example snippets ([#127](https://github.com/Modern-Treasury/modern-treasury-go/issues/127)) ([4b3283b](https://github.com/Modern-Treasury/modern-treasury-go/commit/4b3283bc6bb432f4c39585c8de94896eba56c04c))

## 2.1.0 (2023-10-31)

Full Changelog: [v2.0.2...v2.1.0](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.0.2...v2.1.0)

### Features

* **api:** updates ([#116](https://github.com/Modern-Treasury/modern-treasury-go/issues/116)) ([5b1a909](https://github.com/Modern-Treasury/modern-treasury-go/commit/5b1a9094d01ed28ba41a4ab5beb837c9d2b7a71b))
* **api:** updates ([#121](https://github.com/Modern-Treasury/modern-treasury-go/issues/121)) ([28ed217](https://github.com/Modern-Treasury/modern-treasury-go/commit/28ed217cd027361b68b3f633efb481cae6c0fb23))
* **client:** adjust retry behavior ([#118](https://github.com/Modern-Treasury/modern-treasury-go/issues/118)) ([5cd5611](https://github.com/Modern-Treasury/modern-treasury-go/commit/5cd5611218255d26801ee11e2b0e1fabfac5213f))
* **github:** include a devcontainer setup ([#124](https://github.com/Modern-Treasury/modern-treasury-go/issues/124)) ([25fbe09](https://github.com/Modern-Treasury/modern-treasury-go/commit/25fbe093413bc8c08671061bd0471d8c04b323bf))
* type alias enum values from shared in package root ([#123](https://github.com/Modern-Treasury/modern-treasury-go/issues/123)) ([821b67c](https://github.com/Modern-Treasury/modern-treasury-go/commit/821b67c321582c5e7503bce3f7fc1145624d9192))


### Chores

* **internal:** update gitignore ([#122](https://github.com/Modern-Treasury/modern-treasury-go/issues/122)) ([65eeeab](https://github.com/Modern-Treasury/modern-treasury-go/commit/65eeeab7d51b878f959e40dc5fa0fbc16464bd68))

## 2.0.2 (2023-10-23)

Full Changelog: [v2.0.1...v2.0.2](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.0.1...v2.0.2)

### Bug Fixes

* use /v2 in all package imports ([#114](https://github.com/Modern-Treasury/modern-treasury-go/issues/114)) ([65105b7](https://github.com/Modern-Treasury/modern-treasury-go/commit/65105b7849b435a15c3b41d15bc34313501dc0ac))

## 2.0.1 (2023-10-23)

Full Changelog: [v2.0.0...v2.0.1](https://github.com/Modern-Treasury/modern-treasury-go/compare/v2.0.0...v2.0.1)

### Chores

* **docs:** bump version in README ([#112](https://github.com/Modern-Treasury/modern-treasury-go/issues/112)) ([e05f134](https://github.com/Modern-Treasury/modern-treasury-go/commit/e05f134324347ce6be41ce25e3f132b3058907c2))

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
