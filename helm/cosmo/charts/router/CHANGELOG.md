# Changelog

## [0.7.8](https://github.com/AndreasZeissner/cosmo-ops/compare/helm-cosmo-router@0.7.7...helm-cosmo-router@0.7.8) (2024-08-12)


### Bug Fixes

* test ([2f04683](https://github.com/AndreasZeissner/cosmo-ops/commit/2f04683158d7fbae7ec08d2829e78e9eaeb94429))
* test ([3271c4c](https://github.com/AndreasZeissner/cosmo-ops/commit/3271c4cc86d849bfabe9f877334e606a4955f5d8))
* test ([17bcb1b](https://github.com/AndreasZeissner/cosmo-ops/commit/17bcb1bb55e268ccc09498a5feb63bbdbab6c6f9))
* test ([289aee0](https://github.com/AndreasZeissner/cosmo-ops/commit/289aee02a5b54be078c6aa77183a402571dcaef8))
* test ([bd3a258](https://github.com/AndreasZeissner/cosmo-ops/commit/bd3a2584547ba6834eb55374a044701104b0da63))
* test ([48485e2](https://github.com/AndreasZeissner/cosmo-ops/commit/48485e20e234c717c128f9a1e533f97165a9105d))

## [0.7.7](https://github.com/AndreasZeissner/cosmo-ops/compare/helm-cosmo-router@0.7.6...helm-cosmo-router@0.7.7) (2024-08-12)


### Bug Fixes

* test ([afe87e3](https://github.com/AndreasZeissner/cosmo-ops/commit/afe87e3616bb68deb1c9b51973d08499ab66f266))
* test ([fd4525d](https://github.com/AndreasZeissner/cosmo-ops/commit/fd4525d7be210c08168df9fca347a67b7a119bc6))
* test ([2ec5e47](https://github.com/AndreasZeissner/cosmo-ops/commit/2ec5e47dac33917f218b192023b0b392a937f4dc))
* test ([3d7d47c](https://github.com/AndreasZeissner/cosmo-ops/commit/3d7d47ce768a5acd691a76431212b38c8147fe12))
* test ([e7e0fca](https://github.com/AndreasZeissner/cosmo-ops/commit/e7e0fca8a8293a8f2d0d500eb3f3af0de428c17e))

## [0.7.6](https://github.com/AndreasZeissner/cosmo-ops/compare/helm-cosmo-router@0.7.5...helm-cosmo-router@0.7.6) (2024-08-12)


### Bug Fixes

* test ([fec2abf](https://github.com/AndreasZeissner/cosmo-ops/commit/fec2abfbaba5a993879455b64178dc904f0ef4b2))

## [0.7.5](https://github.com/AndreasZeissner/cosmo-ops/compare/helm-cosmo-router@0.7.4...helm-cosmo-router@0.7.5) (2024-08-12)


### Bug Fixes

* something ([cc13419](https://github.com/AndreasZeissner/cosmo-ops/commit/cc134196190427bc1a2e907e05bf184a276a3e98))

## [0.7.4](https://github.com/AndreasZeissner/cosmo-ops/compare/helm-cosmo-router@0.7.3...helm-cosmo-router@0.7.4) (2024-08-12)


### Bug Fixes

* sthg ([1d8c9c6](https://github.com/AndreasZeissner/cosmo-ops/commit/1d8c9c6afe1a831f4bd5fbf5a05f1a791a39741a))

## [0.7.3](https://github.com/AndreasZeissner/cosmo-ops/compare/helm-cosmo-router@0.7.2...helm-cosmo-router@0.7.3) (2024-08-12)


### Bug Fixes

* router ([16ee75c](https://github.com/AndreasZeissner/cosmo-ops/commit/16ee75c439c06808228ffa32d94d9f1c68bdac47))

## [0.7.2](https://github.com/AndreasZeissner/cosmo-ops/compare/helm-cosmo-router@0.7.1...helm-cosmo-router@0.7.2) (2024-08-12)


### Bug Fixes

* test ([a2b7c97](https://github.com/AndreasZeissner/cosmo-ops/commit/a2b7c97027f7440cc72e4615837644b6bce3cd99))

## [0.7.1](https://github.com/AndreasZeissner/cosmo-ops/compare/helm-cosmo-router@0.7.0...helm-cosmo-router@0.7.1) (2024-08-12)


### Bug Fixes

* note ([25758f8](https://github.com/AndreasZeissner/cosmo-ops/commit/25758f8ecc6ce2a502a3fe087b00f9ecaa4938b2))

## [0.7.0](https://github.com/AndreasZeissner/cosmo-ops/compare/helm-cosmo-router-v0.6.0...helm-cosmo-router@0.7.0) (2024-08-12)


### Features

* add istioGateway and VirtualService to router helm chart ([#773](https://github.com/AndreasZeissner/cosmo-ops/issues/773)) ([2f30950](https://github.com/AndreasZeissner/cosmo-ops/commit/2f30950b1963f8d329bff54c0b7cd8548e4cf207))
* add metrics port to service and deployment spec of router helm chart ([#828](https://github.com/AndreasZeissner/cosmo-ops/issues/828)) ([3e9595b](https://github.com/AndreasZeissner/cosmo-ops/commit/3e9595b3b1ee99c8d7baadf19b42cff8b95a7a43))
* use Helm tpl function on .Values.commonConfiguration ([#840](https://github.com/AndreasZeissner/cosmo-ops/issues/840)) ([acb913e](https://github.com/AndreasZeissner/cosmo-ops/commit/acb913eca34747d6d37a7a84b5c4b188b0e8efa8))


### Bug Fixes

* fix default value for `extraEnvVars` in values for `router` ([#818](https://github.com/AndreasZeissner/cosmo-ops/issues/818)) ([f4cbc28](https://github.com/AndreasZeissner/cosmo-ops/commit/f4cbc28a3a51779eeaaa5108f44e56413caaf005))
* fix the expansion of selectors in the istio-gateway template for the router helm chart ([#785](https://github.com/AndreasZeissner/cosmo-ops/issues/785)) ([07a31a0](https://github.com/AndreasZeissner/cosmo-ops/commit/07a31a07af398c2bfa38be143400a1619aa5876d))

## [0.6.0](https://github.com/wundergraph/cosmo/compare/helm-cosmo-router-v0.5.0...helm-cosmo-router@0.6.0) (2024-08-10)


### Features

* add istioGateway and VirtualService to router helm chart ([#773](https://github.com/wundergraph/cosmo/issues/773)) ([2f30950](https://github.com/wundergraph/cosmo/commit/2f30950b1963f8d329bff54c0b7cd8548e4cf207))
* add metrics port to service and deployment spec of router helm chart ([#828](https://github.com/wundergraph/cosmo/issues/828)) ([3e9595b](https://github.com/wundergraph/cosmo/commit/3e9595b3b1ee99c8d7baadf19b42cff8b95a7a43))
* release helm charts ([#663](https://github.com/wundergraph/cosmo/issues/663)) ([b45c2da](https://github.com/wundergraph/cosmo/commit/b45c2da2a36d7360910eb7c3d2a3207c89d3bbdb))
* use Helm tpl function on .Values.commonConfiguration ([#840](https://github.com/wundergraph/cosmo/issues/840)) ([acb913e](https://github.com/wundergraph/cosmo/commit/acb913eca34747d6d37a7a84b5c4b188b0e8efa8))


### Bug Fixes

* fix default value for `extraEnvVars` in values for `router` ([#818](https://github.com/wundergraph/cosmo/issues/818)) ([f4cbc28](https://github.com/wundergraph/cosmo/commit/f4cbc28a3a51779eeaaa5108f44e56413caaf005))
* fix the expansion of selectors in the istio-gateway template for the router helm chart ([#785](https://github.com/wundergraph/cosmo/issues/785)) ([07a31a0](https://github.com/wundergraph/cosmo/commit/07a31a07af398c2bfa38be143400a1619aa5876d))

## [0.5.0](https://github.com/wundergraph/cosmo/compare/helm-cosmo-router-v0.4.0...helm-cosmo-router@0.5.0) (2024-08-10)


### Features

* add istioGateway and VirtualService to router helm chart ([#773](https://github.com/wundergraph/cosmo/issues/773)) ([2f30950](https://github.com/wundergraph/cosmo/commit/2f30950b1963f8d329bff54c0b7cd8548e4cf207))
* add metrics port to service and deployment spec of router helm chart ([#828](https://github.com/wundergraph/cosmo/issues/828)) ([3e9595b](https://github.com/wundergraph/cosmo/commit/3e9595b3b1ee99c8d7baadf19b42cff8b95a7a43))
* release helm charts ([#663](https://github.com/wundergraph/cosmo/issues/663)) ([b45c2da](https://github.com/wundergraph/cosmo/commit/b45c2da2a36d7360910eb7c3d2a3207c89d3bbdb))
* use Helm tpl function on .Values.commonConfiguration ([#840](https://github.com/wundergraph/cosmo/issues/840)) ([acb913e](https://github.com/wundergraph/cosmo/commit/acb913eca34747d6d37a7a84b5c4b188b0e8efa8))


### Bug Fixes

* fix default value for `extraEnvVars` in values for `router` ([#818](https://github.com/wundergraph/cosmo/issues/818)) ([f4cbc28](https://github.com/wundergraph/cosmo/commit/f4cbc28a3a51779eeaaa5108f44e56413caaf005))
* fix the expansion of selectors in the istio-gateway template for the router helm chart ([#785](https://github.com/wundergraph/cosmo/issues/785)) ([07a31a0](https://github.com/wundergraph/cosmo/commit/07a31a07af398c2bfa38be143400a1619aa5876d))
