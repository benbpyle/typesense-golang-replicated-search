import { Construct } from "constructs";
import { RestApi } from "aws-cdk-lib/aws-apigateway";

export class ApiGatewayConstruct extends Construct {
  private readonly _api: RestApi;

  constructor(scope: Construct, id: string) {
    super(scope, id);

    this._api = new RestApi(this, "RestApi", {
      description: "Typesense Demo API",
      restApiName: "typesense-demo",
      deployOptions: {
        stageName: "main",
      },
    });
  }

  get api(): RestApi {
    return this._api;
  }
}
