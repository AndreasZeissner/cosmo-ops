// @generated by protoc-gen-connect-query v0.6.0 with parameter "target=ts"
// @generated from file wg/cosmo/node/v1/node.proto (package wg.cosmo.node.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetConfigRequest, GetConfigResponse, SelfRegisterRequest, SelfRegisterResponse } from "./node_pb.js";
import { MethodKind } from "@bufbuild/protobuf";
import { createQueryService, createUnaryHooks, UnaryFunctionsWithHooks } from "@connectrpc/connect-query";

export const typeName = "wg.cosmo.node.v1.NodeService";

/**
 * @generated from service wg.cosmo.node.v1.NodeService
 */
export const NodeService = {
  typeName: "wg.cosmo.node.v1.NodeService",
  methods: {
    /**
     * @generated from rpc wg.cosmo.node.v1.NodeService.GetLatestValidRouterConfig
     * @deprecated
     */
    getLatestValidRouterConfig: {
      name: "GetLatestValidRouterConfig",
      I: GetConfigRequest,
      O: GetConfigResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wg.cosmo.node.v1.NodeService.SelfRegister
     */
    selfRegister: {
      name: "SelfRegister",
      I: SelfRegisterRequest,
      O: SelfRegisterResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

const $queryService = createQueryService({  service: NodeService,});

/**
 * @generated from rpc wg.cosmo.node.v1.NodeService.GetLatestValidRouterConfig
 * @deprecated
 */
export const getLatestValidRouterConfig: UnaryFunctionsWithHooks<GetConfigRequest, GetConfigResponse> = {   ...$queryService.getLatestValidRouterConfig,  ...createUnaryHooks($queryService.getLatestValidRouterConfig)};

/**
 * @generated from rpc wg.cosmo.node.v1.NodeService.SelfRegister
 */
export const selfRegister: UnaryFunctionsWithHooks<SelfRegisterRequest, SelfRegisterResponse> = {   ...$queryService.selfRegister,  ...createUnaryHooks($queryService.selfRegister)};
