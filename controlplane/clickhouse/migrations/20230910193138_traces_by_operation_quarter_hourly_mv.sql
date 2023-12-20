-- migrate:up

CREATE MATERIALIZED VIEW IF NOT EXISTS cosmo.traces_by_operation_quarter_hourly_mv TO cosmo.traces_by_operation_quarter_hourly AS SELECT
    toDateTime(
        toStartOfInterval(Timestamp, INTERVAL 15 Minute),
        'UTC'
    ) as Timestamp,
    SpanAttributes [ 'wg.operation.name' ] as OperationName,
    SpanAttributes [ 'wg.operation.type' ] as OperationType,
    SpanAttributes [ 'wg.federated_graph.id'] as FederatedGraphID,
    SpanAttributes ['wg.organization.id'] as OrganizationID,
    mapContains(SpanAttributes, 'wg.subscription') as IsSubscription,
    count() AS TotalRequests,
    countIf(StatusMessage == 'STATUS_CODE_ERROR' OR position(SpanAttributes['http.status_code'],'5') = 1 OR position(SpanAttributes['http.status_code'],'4') = 1 OR SpanAttributes['wg.request.error'] = 'true') AS TotalRequestsError,
    countIf(not(StatusMessage == 'STATUS_CODE_ERROR' OR position(SpanAttributes['http.status_code'],'5') = 1 OR position(SpanAttributes['http.status_code'],'4') = 1 OR SpanAttributes['wg.request.error'] = 'true')) AS TotalRequestsOk,
    quantilesState(0.5, 0.75, 0.9, 0.95, 0.99)(Duration) AS DurationQuantiles,
    toUnixTimestamp(MAX(Timestamp)) as LastCalled
FROM
    cosmo.otel_traces
WHERE empty(ParentSpanId)
GROUP BY
    Timestamp,
    FederatedGraphID,
    OrganizationID,
    OperationName,
    OperationType,
    IsSubscription
ORDER BY
    Timestamp DESC;

-- migrate:down

DROP VIEW IF EXISTS cosmo.traces_by_operation_quarter_hourly_mv