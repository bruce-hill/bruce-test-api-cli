# Changelog

## 0.6.0 (2026-03-06)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/bruce-hill/bruce-test-api-cli/compare/v0.5.0...v0.6.0)

### ⚠ BREAKING CHANGES

* test version bumping
* add support for passing files as parameters

### Features

* add readme documentation for passing files as arguments ([3906583](https://github.com/bruce-hill/bruce-test-api-cli/commit/3906583bef3611f68ffa93f8613d93c6afd0e127))
* add support for file downloads from binary response endpoints ([5fe0119](https://github.com/bruce-hill/bruce-test-api-cli/commit/5fe0119d19fdde72d03525dcd33b0b22890e5835))
* add support for passing files as parameters ([c72e3b8](https://github.com/bruce-hill/bruce-test-api-cli/commit/c72e3b889fa896daae488406d2d7ee6d2cbc42c2))
* **api:** manual updates ([9dc91e1](https://github.com/bruce-hill/bruce-test-api-cli/commit/9dc91e148cdda87fafaa9f69969fa16f4ffc345b))
* **api:** manual updates ([0b56367](https://github.com/bruce-hill/bruce-test-api-cli/commit/0b56367eddf94f0176eec0bbe532087fc03eaa73))
* **api:** manual updates ([2d17c7f](https://github.com/bruce-hill/bruce-test-api-cli/commit/2d17c7fa891e802d8597c9a2e23b9623e3f64a12))
* **api:** manual updates ([4ff0f1d](https://github.com/bruce-hill/bruce-test-api-cli/commit/4ff0f1d6c0c44a52cb8b7181714cb408929b812b))
* **client:** provide file completions when using file embed syntax ([0a041ec](https://github.com/bruce-hill/bruce-test-api-cli/commit/0a041ec0da9f5c030af584ed73ebd6b8bfe52493))
* **cli:** improve shell completions for namespaced commands and flags ([c75c5e2](https://github.com/bruce-hill/bruce-test-api-cli/commit/c75c5e245599ffb36066909fc47ea0b06e4f1546))
* improved documentation and flags for client options ([07ad160](https://github.com/bruce-hill/bruce-test-api-cli/commit/07ad16047d8346c3a1339157d5e47468f2672dbf))
* improved support for passing files for `any`-typed arguments ([076f115](https://github.com/bruce-hill/bruce-test-api-cli/commit/076f1153591f6fc1b2ba35d47c8b0c58fbfc38d6))
* test version bumping ([3b2acd6](https://github.com/bruce-hill/bruce-test-api-cli/commit/3b2acd6cdee5d9417cd7f246df00de146bfd9fff))


### Bug Fixes

* avoid printing usage errors twice ([7850189](https://github.com/bruce-hill/bruce-test-api-cli/commit/78501894d3373c8183ff7a52fcdb3fcc3992b99b))
* fix for file uploads to octet stream and form encoding endpoints ([4a200b8](https://github.com/bruce-hill/bruce-test-api-cli/commit/4a200b82fea8ca00016df6f6876f695e77ec30f3))
* fix for when terminal width is not available ([b8d9d61](https://github.com/bruce-hill/bruce-test-api-cli/commit/b8d9d613b5a0c2126e15efd7a93fed1e066870e8))
* more gracefully handle empty stdin input ([ec9852d](https://github.com/bruce-hill/bruce-test-api-cli/commit/ec9852da81382d42c34240f5cd9d2f7c572a7d49))
* pin formatting for headers to always use repeat/dot formats ([b2f8002](https://github.com/bruce-hill/bruce-test-api-cli/commit/b2f8002774260022dc19f0d635209254acb12c34))
* preserve filename in content-disposition for file uploads ([ca51bb2](https://github.com/bruce-hill/bruce-test-api-cli/commit/ca51bb257ec0b9984ff056cb32c2b2d8c58134d5))
* prevent tests from hanging on streaming/paginated endpoints ([5e88ef2](https://github.com/bruce-hill/bruce-test-api-cli/commit/5e88ef22c1f00768109fa8d338edd590ef77a896))
* remove unused imports ([d3cb3f7](https://github.com/bruce-hill/bruce-test-api-cli/commit/d3cb3f754e8b64f21a4bc892bcf323e97e782cd1))
* use RawJSON for iterated values instead of re-marshalling ([962b6c1](https://github.com/bruce-hill/bruce-test-api-cli/commit/962b6c1d05a9f99f9a25a004c4f5251092ba7224))


### Chores

* add build step to ci ([f5e2b4e](https://github.com/bruce-hill/bruce-test-api-cli/commit/f5e2b4e902d20795f60cdea8d55a0cd0f18056a1))
* **internal:** codegen related update ([6fa526e](https://github.com/bruce-hill/bruce-test-api-cli/commit/6fa526e5571cb834a91a45de966dd25fe3d61874))
* **internal:** codegen related update ([bf28c9f](https://github.com/bruce-hill/bruce-test-api-cli/commit/bf28c9f61ca0444240c6d1603edee25bb033c3ca))
* **internal:** codegen related update ([d91bff7](https://github.com/bruce-hill/bruce-test-api-cli/commit/d91bff79360598980f253a8cc2ea28a052f29742))
* **internal:** codegen related update ([07204f5](https://github.com/bruce-hill/bruce-test-api-cli/commit/07204f5ef9ba644cc0dde8667ed6ab81bfce7eda))
* **internal:** codegen related update ([c509bbb](https://github.com/bruce-hill/bruce-test-api-cli/commit/c509bbb8e3218440ba099ef0c5af5b494a2215af))
* **internal:** codegen related update ([f84e88d](https://github.com/bruce-hill/bruce-test-api-cli/commit/f84e88d8ee1b1fecdd9594ede592b2aa65ee6585))
* **internal:** codegen related update ([953e53d](https://github.com/bruce-hill/bruce-test-api-cli/commit/953e53d14b1e1f2ccc384e08a6fd1ff2ea8323df))
* update documentation in readme ([d2cf254](https://github.com/bruce-hill/bruce-test-api-cli/commit/d2cf2546f7d353f8c189ce08af470778ea3a3648))
* update mock server docs ([98aa7b1](https://github.com/bruce-hill/bruce-test-api-cli/commit/98aa7b1e379fde939fb79e63be82e39f9e95356e))
* update readme with better instructions for installing with homebrew ([7675fed](https://github.com/bruce-hill/bruce-test-api-cli/commit/7675fed3f705dc9db62ba433df0f8d526c454cf2))
* zip READMEs as part of build artifact ([772d0e1](https://github.com/bruce-hill/bruce-test-api-cli/commit/772d0e1277dae110747dcd6dfb18a2431bab014d))

## 0.5.0 (2026-01-26)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/bruce-hill/bruce-test-api-cli/compare/v0.4.0...v0.5.0)

### Features

* **api:** manual updates ([4d69cd3](https://github.com/bruce-hill/bruce-test-api-cli/commit/4d69cd32a598def7cee0e9556af9a307ae3c61b4))
* **api:** manual updates ([1d88bf6](https://github.com/bruce-hill/bruce-test-api-cli/commit/1d88bf6e2d04986af3370e75583bd819e00baee1))
* **api:** manual updates ([b654002](https://github.com/bruce-hill/bruce-test-api-cli/commit/b654002c16ba521e9b70e08b72df6454b0bb9e37))
* **api:** manual updates ([b1c54a4](https://github.com/bruce-hill/bruce-test-api-cli/commit/b1c54a4a6012bb6a769d8eee6ea7dfebba4aa961))
* **api:** manual updates ([dbdea19](https://github.com/bruce-hill/bruce-test-api-cli/commit/dbdea1992778b970f56bbaed99c0be22cf239a7a))
* **api:** manual updates ([be8cf16](https://github.com/bruce-hill/bruce-test-api-cli/commit/be8cf16bce1b79209a182a17083e063fe6d2bb7f))
* **api:** manual updates ([044b0da](https://github.com/bruce-hill/bruce-test-api-cli/commit/044b0da41d6bb8fbe5e46501947e0cf830c4f327))
* **api:** manual updates ([4d1f70c](https://github.com/bruce-hill/bruce-test-api-cli/commit/4d1f70c022f12e65e4b72145b20bd512d78fa12e))
* enable CI tests ([26c6a26](https://github.com/bruce-hill/bruce-test-api-cli/commit/26c6a26f16626d8348c1155a4ee1b3df58a059e6))
* enable suggestion for mistyped commands and flags ([10d15c5](https://github.com/bruce-hill/bruce-test-api-cli/commit/10d15c540e8b01cdb3962c0bb7a76edcf157b580))


### Bug Fixes

* avoid consuming request bodies when printing redacted outputs ([8558ef2](https://github.com/bruce-hill/bruce-test-api-cli/commit/8558ef25d11e65578715b0df0d934f18eaf1718e))
* **client:** do not use pager for short paginated responses ([7ccb35f](https://github.com/bruce-hill/bruce-test-api-cli/commit/7ccb35fa6566963a938da9d45c4c73b39a484f0a))
* fix for paginated output not writing to pager correctly ([070dbe5](https://github.com/bruce-hill/bruce-test-api-cli/commit/070dbe5a74249915de1b7a46ab0d4ad83d813ab9))
* fix terminal height issues causing test failures ([4c190c0](https://github.com/bruce-hill/bruce-test-api-cli/commit/4c190c0b978ff1c034638c8339faeff60963ba07))
* flag types ([65138d3](https://github.com/bruce-hill/bruce-test-api-cli/commit/65138d3083e98f157c42b54e45572882f7a64b55))
* overly broad redaction of Authorization ([8bbc480](https://github.com/bruce-hill/bruce-test-api-cli/commit/8bbc48051beff39fd97b309f12b535c268e73176))
* prevent flag duplication ([ba4c660](https://github.com/bruce-hill/bruce-test-api-cli/commit/ba4c6606ade298708912f211d2a951f858d2e862))
* remove unsupported methods ([cf5d69d](https://github.com/bruce-hill/bruce-test-api-cli/commit/cf5d69d5810147d3dbeffd943a9b80b4414d86fc))


### Chores

* **internal:** codegen related update ([bcb9b15](https://github.com/bruce-hill/bruce-test-api-cli/commit/bcb9b15ae0c1a094cfbed5d6d02cb7342ef3df0d))
* **internal:** codegen related update ([d4bf2bc](https://github.com/bruce-hill/bruce-test-api-cli/commit/d4bf2bcb4db6901226108742646f676f25a5546a))
* **internal:** codegen related update ([a5b89d8](https://github.com/bruce-hill/bruce-test-api-cli/commit/a5b89d80ef54a32fb64d8dd55cba156e54850f68))
* **internal:** codegen related update ([711e405](https://github.com/bruce-hill/bruce-test-api-cli/commit/711e405248a783fc1d22c927126431db782b541b))
* **internal:** codegen related update ([04d4062](https://github.com/bruce-hill/bruce-test-api-cli/commit/04d4062da125e85c2290408840f73ca9b7e92a3e))
* **internal:** codegen related update ([4592687](https://github.com/bruce-hill/bruce-test-api-cli/commit/4592687c177cd2f6612c7354e709f60696501784))
* **internal:** codegen related update ([056eea8](https://github.com/bruce-hill/bruce-test-api-cli/commit/056eea83468b16c14fbffb740dad9dbee16e304f))
* **internal:** update `actions/checkout` version ([189b549](https://github.com/bruce-hill/bruce-test-api-cli/commit/189b5492c39b57dc9645a889ac18b8443b7f2760))
* update Go SDK version ([419a690](https://github.com/bruce-hill/bruce-test-api-cli/commit/419a6908c547d407922e54b72cf2c83d2e31031f))
* update internal comment ([8bbc480](https://github.com/bruce-hill/bruce-test-api-cli/commit/8bbc48051beff39fd97b309f12b535c268e73176))
* updated README.md with more flag information ([f6ac119](https://github.com/bruce-hill/bruce-test-api-cli/commit/f6ac119e968390c5250ebe21ad9dd4798c2bde61))

## 0.4.0 (2026-01-08)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/bruce-hill/bruce-test-api-cli/compare/v0.3.0...v0.4.0)

### Features

* added support for --foo.baz inner field flags ([8b2f21c](https://github.com/bruce-hill/bruce-test-api-cli/commit/8b2f21c6fe91dae105592e73781b4b741c3e666d))
* **api:** manual updates ([885a598](https://github.com/bruce-hill/bruce-test-api-cli/commit/885a598d00ad051089a0d8cf0bd9e127789402dd))
* **api:** manual updates ([6f126e3](https://github.com/bruce-hill/bruce-test-api-cli/commit/6f126e306f532b0343368ffa33240db8b9cd8023))
* **api:** manual updates ([2b8ffe3](https://github.com/bruce-hill/bruce-test-api-cli/commit/2b8ffe3e3727cf9b7926dbe6ecf61cebab1e079c))
* **api:** manual updates ([e1b1ef2](https://github.com/bruce-hill/bruce-test-api-cli/commit/e1b1ef234119696e45b19e7a4f048333cd1abaa6))
* **api:** manual updates ([f6aa297](https://github.com/bruce-hill/bruce-test-api-cli/commit/f6aa297742298dbe74dec8bd3a30e12167455b47))
* **api:** manual updates ([369ed2c](https://github.com/bruce-hill/bruce-test-api-cli/commit/369ed2c2b44e0c2e4f7023ba388b2822081310ed))
* **api:** manual updates ([56d42d8](https://github.com/bruce-hill/bruce-test-api-cli/commit/56d42d8b5295ed12d39950a408fafe8a444ff648))
* **api:** manual updates ([313b153](https://github.com/bruce-hill/bruce-test-api-cli/commit/313b1532b8dad26a2c6d672c24bc478d42aeea86))
* **api:** manual updates ([095aebf](https://github.com/bruce-hill/bruce-test-api-cli/commit/095aebf061bd27ddca9e2e23af63ea0cf7e92218))
* **api:** manual updates ([2b1e5a1](https://github.com/bruce-hill/bruce-test-api-cli/commit/2b1e5a1a72554a76669b3cf84875242e5de013da))
* **api:** manual updates ([3722c55](https://github.com/bruce-hill/bruce-test-api-cli/commit/3722c5595f19f6102eca74d82337642be5a8ec8a))
* improved behavior for exploring paginated/streamed endpoints ([a6d3983](https://github.com/bruce-hill/bruce-test-api-cli/commit/a6d39834faf6a200d02650f1ab2496e3b655feeb))


### Bug Fixes

* check required arguments ([d67dacc](https://github.com/bruce-hill/bruce-test-api-cli/commit/d67dacc8dedcab06876f3b1a961e02b434795d8f))
* fixed placeholders for date/time arguments ([d2b47fa](https://github.com/bruce-hill/bruce-test-api-cli/commit/d2b47fac04174195a01f4500c8c9cdc17258ee33))

## 0.3.0 (2025-12-19)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/bruce-hill/bruce-test-api-cli/compare/v0.2.0...v0.3.0)

### ⚠ BREAKING CHANGES

* new logic for parsing arguments

### Features

* add better suggests when commands don't match ([f94ae2d](https://github.com/bruce-hill/bruce-test-api-cli/commit/f94ae2d7a4f0957e0fa14ddca494668f0c0eb18c))
* added `--output-filter` flag and `--error-format` flag to support better visualization options ([32ab7a4](https://github.com/bruce-hill/bruce-test-api-cli/commit/32ab7a4e82db8fae569eddad16aa38c77265f18e))
* added mock server tests ([7c27f5c](https://github.com/bruce-hill/bruce-test-api-cli/commit/7c27f5c114d26658a7dbf0e4d8e90928ff375157))
* **api:** descriptions ([2daa2f4](https://github.com/bruce-hill/bruce-test-api-cli/commit/2daa2f47f37cfd1ac856122eb33beebbba98b08c))
* **api:** manual updates ([1dbdc68](https://github.com/bruce-hill/bruce-test-api-cli/commit/1dbdc682f702301f49d23139a9785ca4da480be0))
* **api:** manual updates ([b3c7c46](https://github.com/bruce-hill/bruce-test-api-cli/commit/b3c7c46eb7ab933d6235924286779ed1b172e1fc))
* **api:** manual updates ([9b3ba98](https://github.com/bruce-hill/bruce-test-api-cli/commit/9b3ba98aa8819f3c0ecdff776739fd884b9c442d))
* **api:** manual updates ([8ad904f](https://github.com/bruce-hill/bruce-test-api-cli/commit/8ad904fbbfdc63194208d7b46ab80135efd81275))
* **api:** manual updates ([ad89442](https://github.com/bruce-hill/bruce-test-api-cli/commit/ad8944266e335b4f25faabc77a358a744db8563e))
* **api:** manual updates ([8079ca7](https://github.com/bruce-hill/bruce-test-api-cli/commit/8079ca7a8ebaf99984712f6130df25e921cd29c2))
* **api:** manual updates ([d142428](https://github.com/bruce-hill/bruce-test-api-cli/commit/d142428b58bb52f8579cef126309c3b413cd7897))
* **api:** manual updates ([9814cfb](https://github.com/bruce-hill/bruce-test-api-cli/commit/9814cfb77ec3efff7cd7fb0a61782e42d6b2db6a))
* **api:** manual updates ([c83b5ff](https://github.com/bruce-hill/bruce-test-api-cli/commit/c83b5ff22751417c77c303afc0ac68a85eb0d123))
* **api:** manual updates ([7dbaa62](https://github.com/bruce-hill/bruce-test-api-cli/commit/7dbaa622c17fed7cb2c8b569a429f5565097f25f))
* **api:** manual updates ([63ab7ca](https://github.com/bruce-hill/bruce-test-api-cli/commit/63ab7caeaf027991425485a3596c9c0bf3dc3951))
* **api:** manual updates ([18a65ed](https://github.com/bruce-hill/bruce-test-api-cli/commit/18a65ed85abd5bb00423c44f47b01bdc7b0c1dc7))
* **api:** manual updates ([2f12a67](https://github.com/bruce-hill/bruce-test-api-cli/commit/2f12a67e14191a7404e0ea73f6255d5956464e49))
* **api:** manual updates ([2c6c20c](https://github.com/bruce-hill/bruce-test-api-cli/commit/2c6c20cb019e836e05b0a6783f8ab367fe4c3e07))
* **api:** manual updates ([ce03db8](https://github.com/bruce-hill/bruce-test-api-cli/commit/ce03db834a2896f59807c15c27cbf1d01d1be6bb))
* **api:** manual updates ([df12411](https://github.com/bruce-hill/bruce-test-api-cli/commit/df12411dbf594ba8d7c8595440bb5c0d8d579daa))
* **api:** manual updates ([0c2d955](https://github.com/bruce-hill/bruce-test-api-cli/commit/0c2d955575abc66e3c65d5bab21ccbd46bd803dd))
* **api:** manual updates ([967288d](https://github.com/bruce-hill/bruce-test-api-cli/commit/967288d795ba8c5d251e558b72e55f96743f29bc))
* **api:** manual updates ([73d80a2](https://github.com/bruce-hill/bruce-test-api-cli/commit/73d80a277763d0489f2b8955fb704d814b160d42))
* **api:** manual updates ([81f9768](https://github.com/bruce-hill/bruce-test-api-cli/commit/81f9768c9dd747fbc31c7ea5271a31c2bca5cd44))
* **api:** manual updates ([abd6c09](https://github.com/bruce-hill/bruce-test-api-cli/commit/abd6c09cf3c6a0bd9aafbc2c1741027cc2c9ab21))
* **api:** manual updates ([12cb3fe](https://github.com/bruce-hill/bruce-test-api-cli/commit/12cb3fedbdea35b2ddf99959c64f2f181f1de97c))
* **api:** manual updates ([b4fe9c2](https://github.com/bruce-hill/bruce-test-api-cli/commit/b4fe9c202d841b4696939b995bdb23e8313d8359))
* **api:** manual updates ([e15c126](https://github.com/bruce-hill/bruce-test-api-cli/commit/e15c12679b1d7f8577fd15c94964d3480dabaad6))
* **api:** manual updates ([cc4b8c2](https://github.com/bruce-hill/bruce-test-api-cli/commit/cc4b8c2dcedf77093c3bffee6a78c2325f7aea53))
* **api:** manual updates ([c2b82da](https://github.com/bruce-hill/bruce-test-api-cli/commit/c2b82dad153e62ccad56a4dbb5d04b0dc7e811c5))
* **api:** manual updates ([650841b](https://github.com/bruce-hill/bruce-test-api-cli/commit/650841b4a2c9058642666ff594b96cbf2a350f7f))
* **api:** manual updates ([68e7753](https://github.com/bruce-hill/bruce-test-api-cli/commit/68e775303044bc0b40eb9568840e52206f31811f))
* **api:** manual updates ([2ec4d8d](https://github.com/bruce-hill/bruce-test-api-cli/commit/2ec4d8d51232157efe04b6b5b57cef2ac0253812))
* **api:** manual updates ([2f53781](https://github.com/bruce-hill/bruce-test-api-cli/commit/2f5378126cf986874017dbc35a5ad5cf6c0a4d90))
* **api:** manual updates ([e1def37](https://github.com/bruce-hill/bruce-test-api-cli/commit/e1def377541a4921aaf738622010818740955fb7))
* **api:** manual updates ([77e8df3](https://github.com/bruce-hill/bruce-test-api-cli/commit/77e8df3a395d558e5f75130c6eb07a3b6150bf88))
* **api:** manual updates ([a48ed0b](https://github.com/bruce-hill/bruce-test-api-cli/commit/a48ed0bc248e908b151a7b5a89b22a9d9787d798))
* **api:** manual updates ([dc6a33d](https://github.com/bruce-hill/bruce-test-api-cli/commit/dc6a33dbcf6ecebc329e89a97051f76c036720e2))
* **api:** manual updates ([a7e9b92](https://github.com/bruce-hill/bruce-test-api-cli/commit/a7e9b9264a1363771e2a9bdfd02df178040f405f))
* **api:** manual updates ([ebc604e](https://github.com/bruce-hill/bruce-test-api-cli/commit/ebc604e8ed1ba08d1bb8b1e6fc750668ff4d4cf5))
* **api:** manual updates ([65e25ae](https://github.com/bruce-hill/bruce-test-api-cli/commit/65e25ae6273e8df5a86300ae76c740d8976775b2))
* **api:** manual updates ([9174eed](https://github.com/bruce-hill/bruce-test-api-cli/commit/9174eed312ee500d915446a73157335a0992dc37))
* **api:** manual updates ([6148021](https://github.com/bruce-hill/bruce-test-api-cli/commit/6148021091e773cf96c03163d6d1d9ffb5fabe98))
* **api:** manual updates ([d19df11](https://github.com/bruce-hill/bruce-test-api-cli/commit/d19df11987e1d20773c727e3dc583ce9204b70d8))
* **api:** manual updates ([e38a2aa](https://github.com/bruce-hill/bruce-test-api-cli/commit/e38a2aa63585422d6929ac263f8705d41947959f))
* **api:** manual updates ([76e8b0d](https://github.com/bruce-hill/bruce-test-api-cli/commit/76e8b0d93f48efea0f3859a7b548aa9ac5715fe7))
* **api:** manual updates ([fc137ee](https://github.com/bruce-hill/bruce-test-api-cli/commit/fc137eec6ebccf1e4d0873517d3331ee982bb0ce))
* **api:** manual updates ([1258424](https://github.com/bruce-hill/bruce-test-api-cli/commit/12584241beb45415ad389c2f5877cfa8a2957fa3))
* **api:** manual updates ([64dd4c4](https://github.com/bruce-hill/bruce-test-api-cli/commit/64dd4c4b1690d60204faba310425ab79296acbc7))
* **api:** manual updates ([c174e05](https://github.com/bruce-hill/bruce-test-api-cli/commit/c174e05bd4b3c161bec1bff8871ef8f15b2c615d))
* **api:** manual updates ([e09b816](https://github.com/bruce-hill/bruce-test-api-cli/commit/e09b816a7c610b393c805ee367f2a8103d1f3497))
* **api:** manual updates ([3027810](https://github.com/bruce-hill/bruce-test-api-cli/commit/3027810106830f6d0b69ae7ed73347888774a819))
* **api:** manual updates ([352fa34](https://github.com/bruce-hill/bruce-test-api-cli/commit/352fa346c26fd73a138b8f346f771eac0eac475e))
* **api:** manual updates ([51a0673](https://github.com/bruce-hill/bruce-test-api-cli/commit/51a0673254874871a12a977029c831c17289212f))
* **api:** manual updates ([9d3d35d](https://github.com/bruce-hill/bruce-test-api-cli/commit/9d3d35dd846c218c85354d4cf8404cb262b1e5ae))
* **api:** manual updates ([8eec888](https://github.com/bruce-hill/bruce-test-api-cli/commit/8eec8883b2b93285228992401758ab1b989353a9))
* **api:** manual updates ([bd2070d](https://github.com/bruce-hill/bruce-test-api-cli/commit/bd2070df1002440bb00212cb419a98ae86e6a541))
* **api:** manual updates ([3cc78ac](https://github.com/bruce-hill/bruce-test-api-cli/commit/3cc78ac17eb6ea5f5321bb0599053bc68d85ebb7))
* **api:** people/pets ([e97eb9b](https://github.com/bruce-hill/bruce-test-api-cli/commit/e97eb9b58175ad6fa67f9fa7e546efb40d723e4a))
* **api:** update via SDK Studio ([3a90c98](https://github.com/bruce-hill/bruce-test-api-cli/commit/3a90c98e80a5fa2995e55cb7f0d917152e6d3be4))
* arguments now have defaults and descriptions ([51cbf84](https://github.com/bruce-hill/bruce-test-api-cli/commit/51cbf84dc55952f9276a8161071d2133a0ed108a))
* better support for positional arguments ([bc6bc26](https://github.com/bruce-hill/bruce-test-api-cli/commit/bc6bc26bdce881cfd6da2519e61c391e63944af7))
* **cli:** automatic streaming for paginated endpoints ([d375e1a](https://github.com/bruce-hill/bruce-test-api-cli/commit/d375e1a93536ebe35473548eb32dd1a52ded1949))
* **cli:** binary request bodies ([4e202f2](https://github.com/bruce-hill/bruce-test-api-cli/commit/4e202f228415e708720bb11169eb1c1a8eb8163e))
* enable support for streaming endpoints ([222acbd](https://github.com/bruce-hill/bruce-test-api-cli/commit/222acbd2922f7302d258e9367c2301b7a342b856))
* fix edge cases for sending request data and add YAML support ([8f11920](https://github.com/bruce-hill/bruce-test-api-cli/commit/8f11920b8821ce540d71ec4e8057ec7890aeffa6))
* new logic for parsing arguments ([ede380d](https://github.com/bruce-hill/bruce-test-api-cli/commit/ede380db4d7307cc17dbbd16afcb8a1e360ac07a))
* now ships with manpages ([c8c3c6d](https://github.com/bruce-hill/bruce-test-api-cli/commit/c8c3c6d34b6e73ac296e8c6610219f57e6436ed8))
* redact `Authorization` header when using debug option ([c4a01a0](https://github.com/bruce-hill/bruce-test-api-cli/commit/c4a01a0088cdeafeed9090398523cbfe01396c3d))


### Bug Fixes

* **cli:** fix compilation on Windows ([57f4f5c](https://github.com/bruce-hill/bruce-test-api-cli/commit/57f4f5c6a442affeb8883f0d7336926e64391b80))
* **cli:** remove `*.exe` files from customer SDK changes ([4114fd1](https://github.com/bruce-hill/bruce-test-api-cli/commit/4114fd181010f9c1d787aee7b45bff47b834ff5e))
* **cli:** when encoutering scalar body root parameter, don't treat it as binary ([e7e3d4d](https://github.com/bruce-hill/bruce-test-api-cli/commit/e7e3d4d82093cc847d2466000d4064e6048fa1c1))
* downgrade urfave/cli-docs dependency ([9249509](https://github.com/bruce-hill/bruce-test-api-cli/commit/9249509aec379962c443e39b3e2065dc0a95e0bf))
* fix builds for non-public Go repos ([a53744a](https://github.com/bruce-hill/bruce-test-api-cli/commit/a53744a483285ba994b180cc610c46ae1178cdba))
* fix for default flag values ([5f8a7a4](https://github.com/bruce-hill/bruce-test-api-cli/commit/5f8a7a452a5fb1fdf7ebdcbfb64320ceb627dbe8))
* fix for empty request bodies ([a44a85f](https://github.com/bruce-hill/bruce-test-api-cli/commit/a44a85fe1a3ce7def843bc723161ffa54a030423))
* fix for issue with nil responses ([3a024dc](https://github.com/bruce-hill/bruce-test-api-cli/commit/3a024dcedc9b5edd034ed050cd4100a604aaa542))
* fix paginated endpoints for primitive types ([364f834](https://github.com/bruce-hill/bruce-test-api-cli/commit/364f8345faa38b4da483e4d70163cc1dfebb0a5a))
* fixed manpage generation ([c418944](https://github.com/bruce-hill/bruce-test-api-cli/commit/c4189441f8ae17e2d1a280340fa0f01f61c05740))
* **homebrew:** homebrew distribution should work now ([8990114](https://github.com/bruce-hill/bruce-test-api-cli/commit/89901145d94e59d459d86a9aab24b8cecbcb732e))
* ignore .exe files ([1013117](https://github.com/bruce-hill/bruce-test-api-cli/commit/10131173967d50935036cb431981a185381016b2))
* **mcp:** correct code tool API endpoint ([21d877c](https://github.com/bruce-hill/bruce-test-api-cli/commit/21d877c5e108fb126b2781483ba4bcc5e0c69525))
* pass through context parameter correctly ([7ba493e](https://github.com/bruce-hill/bruce-test-api-cli/commit/7ba493e5c3bd9d365c1e480027446c7c8f15bdc0))
* remove some bootstrapping logic ([5eadc69](https://github.com/bruce-hill/bruce-test-api-cli/commit/5eadc6965ff89a7c69d095e56762d8271755c3bb))


### Chores

* add scripts ([49796dc](https://github.com/bruce-hill/bruce-test-api-cli/commit/49796dcb9435c0922a98267d37c1e0ced104bb04))
* bump Go version ([4d55421](https://github.com/bruce-hill/bruce-test-api-cli/commit/4d55421b1cf7e013b3a9620b6b0cd60d8a11fcaa))
* **cli:** add `*.exe` files back to `.gitignore` ([163924e](https://github.com/bruce-hill/bruce-test-api-cli/commit/163924e8fcca68d84026a03b59f129e6b672da7a))
* **cli:** move `jsonview` subpackage to `internal` ([e5d6179](https://github.com/bruce-hill/bruce-test-api-cli/commit/e5d6179cd2dd3b691e0362ab7de70c395cabd338))
* **cli:** run pre-codegen tests on Windows ([8b9a336](https://github.com/bruce-hill/bruce-test-api-cli/commit/8b9a33619e37a07e5f7fee72e80b5220c1895774))
* **cli:** temporarily remove `*.exe` from `.gitignore` ([1a30213](https://github.com/bruce-hill/bruce-test-api-cli/commit/1a30213581efb9375e8462dae981b68229590d5a))
* code cleanup for `interface{}` ([014dbbc](https://github.com/bruce-hill/bruce-test-api-cli/commit/014dbbcefaccd7539b7d82bc70e36e788d4258be))
* configure new SDK language ([e17255d](https://github.com/bruce-hill/bruce-test-api-cli/commit/e17255d985ff5a3028c3bd26d28973e24cbed24d))
* do not install brew dependencies in ./scripts/bootstrap by default ([ebd69c6](https://github.com/bruce-hill/bruce-test-api-cli/commit/ebd69c66fbd7495e48283ed3452ea130cca7a14c))
* **internal:** codegen related update ([e14e304](https://github.com/bruce-hill/bruce-test-api-cli/commit/e14e304c4a65e8857a55c573dd4c76c0cc6b38f4))
* **internal:** codegen related update ([4f0272a](https://github.com/bruce-hill/bruce-test-api-cli/commit/4f0272a1ee253f31ec62bb24fa75290fe4ebaf4e))
* **internal:** codegen related update ([4856cd6](https://github.com/bruce-hill/bruce-test-api-cli/commit/4856cd665edb0d18eff9f9b08509ae5a34f626f7))
* **internal:** codegen related update ([b2850d3](https://github.com/bruce-hill/bruce-test-api-cli/commit/b2850d33a58d1ebf1d0f1d2674ad9599859c8f62))
* **internal:** codegen related update ([e784840](https://github.com/bruce-hill/bruce-test-api-cli/commit/e7848409e06598d52800d5425471f4acc354619a))
* **internal:** codegen related update ([c4594ab](https://github.com/bruce-hill/bruce-test-api-cli/commit/c4594ab6ae3aca74e50289a23f481359c17fdf72))
* **internal:** codegen related update ([3ba560a](https://github.com/bruce-hill/bruce-test-api-cli/commit/3ba560a22464571ae72dad5fd380a1b7baab46ae))
* **internal:** codegen related update ([cbd3dfc](https://github.com/bruce-hill/bruce-test-api-cli/commit/cbd3dfc970d6deee341ccd98e020d991419a5c5d))
* **internal:** codegen related update ([364a22b](https://github.com/bruce-hill/bruce-test-api-cli/commit/364a22bae549f6134e8131c78aaf826cf8888147))
* **internal:** codegen related update ([38a736a](https://github.com/bruce-hill/bruce-test-api-cli/commit/38a736ac4f5d280f051fe0d564d492ad0e4b99f3))
* **internal:** codegen related update ([1b2453e](https://github.com/bruce-hill/bruce-test-api-cli/commit/1b2453e1b2707a0ee38ba41801774f514f3f1e27))
* **internal:** codegen related update ([cd3b4ce](https://github.com/bruce-hill/bruce-test-api-cli/commit/cd3b4ce2a3dc2721ae7fc90cb246b22494429561))
* **internal:** codegen related update ([4fcbd6c](https://github.com/bruce-hill/bruce-test-api-cli/commit/4fcbd6c2e89003c36443b860ce2a425745dd7e7b))
* **internal:** codegen related update ([b40e24b](https://github.com/bruce-hill/bruce-test-api-cli/commit/b40e24bf810bc11781c8d9e2aa30981b92d98ff3))
* **internal:** codegen related update ([81ecc63](https://github.com/bruce-hill/bruce-test-api-cli/commit/81ecc63990194cc2d71acebe4473a66dcdb09884))
* **internal:** codegen related update ([313b13c](https://github.com/bruce-hill/bruce-test-api-cli/commit/313b13cbc1e37ab7387076a35143721b9b72d1d4))
* **internal:** codegen related update ([2f1bdb6](https://github.com/bruce-hill/bruce-test-api-cli/commit/2f1bdb66a2468c89d31dce21fdbce191823c222c))
* **internal:** codegen related update ([f63c354](https://github.com/bruce-hill/bruce-test-api-cli/commit/f63c354760dc56c85d29c5c943393974d9cb94df))
* **internal:** codegen related update ([bc875ba](https://github.com/bruce-hill/bruce-test-api-cli/commit/bc875ba2d25dfece12d31d45da6cb5daec330002))
* **internal:** codegen related update ([86ef8f8](https://github.com/bruce-hill/bruce-test-api-cli/commit/86ef8f818f68e82b91dbcc11728f6e5450dbcbbd))
* **internal:** codegen related update ([36964b9](https://github.com/bruce-hill/bruce-test-api-cli/commit/36964b9ee46de27213e23465458d8eb7e04dc329))
* **internal:** codegen related update ([67080b1](https://github.com/bruce-hill/bruce-test-api-cli/commit/67080b16f4a3e1549a3c7f0ec331217e5de62de9))
* **internal:** codegen related update ([b877c36](https://github.com/bruce-hill/bruce-test-api-cli/commit/b877c36c703d53c6e767a8356ba915a476416812))
* **internal:** codegen related update ([3b22666](https://github.com/bruce-hill/bruce-test-api-cli/commit/3b2266694073136413af3e34199593e1157ae77d))
* sync repo ([87203a6](https://github.com/bruce-hill/bruce-test-api-cli/commit/87203a61732599b9ab33a5f214fe24cecc40e588))
* update dependencies ([ac936cc](https://github.com/bruce-hill/bruce-test-api-cli/commit/ac936ccb7eae2cc019c530ac881c9934cd67b2a1))
* use `stretchr/testify` assertion helpers in tests ([71bca99](https://github.com/bruce-hill/bruce-test-api-cli/commit/71bca99d60b77691b47eb867ba8640d535b93e1a))

## 0.2.0 (2025-09-03)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/bruce-hill/bruce-test-api-cli/compare/v0.1.0...v0.2.0)

### Features

* **api:** manual updates ([4ec4519](https://github.com/bruce-hill/bruce-test-api-cli/commit/4ec4519591dba5e93b250abb3e79b978257a5267))
* **api:** manual updates ([b0c2e98](https://github.com/bruce-hill/bruce-test-api-cli/commit/b0c2e980bba6cad72ee4ca2a2ee59ef1d605e301))
* **api:** manual updates ([c5ad8ca](https://github.com/bruce-hill/bruce-test-api-cli/commit/c5ad8ca56111e4c93a8f4f7bdecea4eb8822198c))
* **api:** manual updates ([993b9ef](https://github.com/bruce-hill/bruce-test-api-cli/commit/993b9efcff83690def572ff5e9fd1c3bdb65c918))
* **api:** manual updates ([69f9396](https://github.com/bruce-hill/bruce-test-api-cli/commit/69f93968cf21f971afd17da5ee3ceadd17fc7945))
* **api:** manual updates ([b4c8486](https://github.com/bruce-hill/bruce-test-api-cli/commit/b4c8486f7f815c7c730b32ec68128c225b0363ae))
* **api:** manual updates ([534f368](https://github.com/bruce-hill/bruce-test-api-cli/commit/534f368be36b01a0c299bf1bd9cfb67f7d812fa8))
* **api:** manual updates ([bdda8b8](https://github.com/bruce-hill/bruce-test-api-cli/commit/bdda8b858feb5301744d082e9a3b2b07c74cbf61))
* **api:** manual updates ([45b9363](https://github.com/bruce-hill/bruce-test-api-cli/commit/45b936368816b990164bdb61d9bf50139b8cbaba))

## 0.1.0 (2025-09-02)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/bruce-hill/bruce-test-api-cli/compare/v0.0.1...v0.1.0)

### Features

* **api:** manual updates ([c94d965](https://github.com/bruce-hill/bruce-test-api-cli/commit/c94d9654c4c6c6e90e83bf0c592acaf14b30852a))
* **api:** manual updates ([6409f08](https://github.com/bruce-hill/bruce-test-api-cli/commit/6409f08f46d5710620cbc7dcd24027e37c6dafc2))
* **api:** manual updates ([cd7c10d](https://github.com/bruce-hill/bruce-test-api-cli/commit/cd7c10da090f407f95aa7fd0bd39f22b704c6850))


### Chores

* update SDK settings ([84482d9](https://github.com/bruce-hill/bruce-test-api-cli/commit/84482d9eb8b477a0e0f42626c67e12d23b45f5f6))
* update SDK settings ([c5e81f7](https://github.com/bruce-hill/bruce-test-api-cli/commit/c5e81f7c0de22aff61def5d4a998a7e45f1c931a))
