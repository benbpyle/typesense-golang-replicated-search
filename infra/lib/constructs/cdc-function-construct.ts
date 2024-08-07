import { Construct } from "constructs";
import { GoFunction } from "@aws-cdk/aws-lambda-go-alpha";
import { Table } from "aws-cdk-lib/aws-dynamodb";
import { Duration } from "aws-cdk-lib";
import * as path from "path";
import { DynamoEventSource } from "aws-cdk-lib/aws-lambda-event-sources";
import { StartingPosition } from "aws-cdk-lib/aws-lambda";

export default class CdcFunctionConstruct extends Construct {
  constructor(scope: Construct, id: string, table: Table) {
    super(scope, id);

    const func = new GoFunction(scope, "CdcFunction", {
      entry: path.join(__dirname, `../../../src/cdc-function`),
      functionName: "typesense-demo-cdc-function",
      timeout: Duration.seconds(30),
      environment: {
        LOG_LEVEL: "debug",
        TABLE_NAME: table.tableName,
      },
    });

    func.addEventSource(
      new DynamoEventSource(table, {
        startingPosition: StartingPosition.LATEST,
      }),
    );
    table.grantStream(func);
  }
}
