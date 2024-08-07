import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { TableConstruct } from "./constructs/table-construct";
import PostFunctionConstruct from "./constructs/post-function-construct";
import { ApiGatewayConstruct } from "./constructs/api-gateway-construct";
import CdcFunctionConstruct from "./constructs/cdc-function-construct";
import SearchFunctionConstruct from "./constructs/search-function-construct";
// import * as sqs from 'aws-cdk-lib/aws-sqs';

export class GolangTypesenseStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const tableConstruct = new TableConstruct(this, "TableConstruct");
    const apiGatewayConstruct = new ApiGatewayConstruct(
      this,
      "ApiGatewayConstruct",
    );
    const postFunctionConstruct = new PostFunctionConstruct(
      this,
      "PostFunctionConstruct",
      apiGatewayConstruct.api,
      tableConstruct.table,
    );

    const cdcFunctionConstruct = new CdcFunctionConstruct(
      this,
      "CdcFunctionConstruct",
      tableConstruct.table,
    );

    const searchFunctionConstruct = new SearchFunctionConstruct(
      this,
      "SearchFunctionConstruct",
      apiGatewayConstruct.api,
    );
  }
}
