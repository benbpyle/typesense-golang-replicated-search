import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import CdcFunctionConstruct from "./constructs/cdc-function-construct";

export class GolangTypesenseStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const cdcFunctionConstruct = new CdcFunctionConstruct(
      this,
      "CdcFunctionConstruct",
      // tableConstruct.table,
    );
  }
}
