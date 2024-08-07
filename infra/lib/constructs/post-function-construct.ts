import {
  LambdaIntegration,
  Method,
  Resource,
  RestApi,
} from "aws-cdk-lib/aws-apigateway";
import { Construct } from "constructs";
import { GoFunction } from "@aws-cdk/aws-lambda-go-alpha";
import { Table } from "aws-cdk-lib/aws-dynamodb";
import { Duration } from "aws-cdk-lib";
import * as path from "path";

export default class PostFunctionConstruct extends Construct {
  constructor(scope: Construct, id: string, api: RestApi, table: Table) {
    super(scope, id);

    const func = new GoFunction(scope, "PostFunction", {
      entry: path.join(__dirname, `../../../src/post-function`),
      functionName: "typesense-demo-post-function",
      timeout: Duration.seconds(30),
      environment: {
        LOG_LEVEL: "debug",
        TABLE_NAME: table.tableName,
      },
    });

    api.root.addMethod(
      "POST",
      new LambdaIntegration(func, {
        proxy: true,
      }),
    );

    table.grantReadWriteData(func);
  }
}
