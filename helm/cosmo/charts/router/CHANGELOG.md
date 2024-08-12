# Changelog

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
