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

export default class SearchFunctionConstruct extends Construct {
  constructor(scope: Construct, id: string, api: RestApi) {
    super(scope, id);

    const func = new GoFunction(scope, "SearchFunction", {
      entry: path.join(__dirname, `../../../src/search-function`),
      functionName: "typesense-demo-search-function",
      timeout: Duration.seconds(30),
      environment: {
        LOG_LEVEL: "debug",
        TYPESENSE_CLUSTER_URL: process.env.TYPESENSE_CLUSTER_URL!,
        TYPESENSE_API_KEY: process.env.TYPESENSE_API_KEY!,
      },
    });

    const resource = new Resource(scope, "SearchResource", {
      parent: api.root,
      pathPart: "search",
    });

    resource.addMethod(
      "GET",
      new LambdaIntegration(func, {
        proxy: true,
      }),
    );
  }
}
