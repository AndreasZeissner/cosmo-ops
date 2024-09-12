# Changelog

## [0.9.0](https://github.com/AndreasZeissner/cosmo-ops/compare/helm-cosmo-router-v0.8.0...helm-cosmo-router@0.9.0) (2024-09-12)


### Features

* **cosmo:** add setting commonLabels on subchart resources ([#1120](https://github.com/AndreasZeissner/cosmo-ops/issues/1120)) ([030dc6d](https://github.com/AndreasZeissner/cosmo-ops/commit/030dc6da6652508d041bb34715d867d3a54db004))
* enable using HTTP(S)_PROXY in router  ([#1136](https://github.com/AndreasZeissner/cosmo-ops/issues/1136)) ([4600fdf](https://github.com/AndreasZeissner/cosmo-ops/commit/4600fdff6ab57541a6119e4e51180ed4403363a6))


### Bug Fixes

* test ([2f04683](https://github.com/AndreasZeissner/cosmo-ops/commit/2f04683158d7fbae7ec08d2829e78e9eaeb94429))
* test ([3271c4c](https://github.com/AndreasZeissner/cosmo-ops/commit/3271c4cc86d849bfabe9f877334e606a4955f5d8))
* test ([17bcb1b](https://github.com/AndreasZeissner/cosmo-ops/commit/17bcb1bb55e268ccc09498a5feb63bbdbab6c6f9))
* test ([289aee0](https://github.com/AndreasZeissner/cosmo-ops/commit/289aee02a5b54be078c6aa77183a402571dcaef8))
* test ([bd3a258](https://github.com/AndreasZeissner/cosmo-ops/commit/bd3a2584547ba6834eb55374a044701104b0da63))
* test ([48485e2](https://github.com/AndreasZeissner/cosmo-ops/commit/48485e20e234c717c128f9a1e533f97165a9105d))

## [0.8.0](https://github.com/wundergraph/cosmo/compare/helm-cosmo-router@0.7.0...helm-cosmo-router@0.8.0) (2024-09-05)


### Features

* enable using HTTP(S)_PROXY in router  ([#1136](https://github.com/wundergraph/cosmo/issues/1136)) ([4600fdf](https://github.com/wundergraph/cosmo/commit/4600fdff6ab57541a6119e4e51180ed4403363a6))

## [0.7.0](https://github.com/wundergraph/cosmo/compare/helm-cosmo-router@0.6.0...helm-cosmo-router@0.7.0) (2024-08-29)


### Features

* **cosmo:** add setting commonLabels on subchart resources ([#1120](https://github.com/wundergraph/cosmo/issues/1120)) ([030dc6d](https://github.com/wundergraph/cosmo/commit/030dc6da6652508d041bb34715d867d3a54db004))

## [0.6.0](https://github.com/wundergraph/cosmo/compare/helm-cosmo-router-v0.5.0...helm-cosmo-router@0.6.0) (2024-08-14)


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
