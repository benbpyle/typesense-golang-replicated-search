import { RemovalPolicy } from "aws-cdk-lib";
import * as db from "aws-cdk-lib/aws-dynamodb";
import { Table } from "aws-cdk-lib/aws-dynamodb";
import { Construct } from "constructs";

export class TableConstruct extends Construct {
  private _table: Table;

  constructor(scope: Construct, id: string) {
    super(scope, id);
    this.buildTable(scope);
  }

  get table(): Table {
    return this._table;
  }

  private buildTable = (scope: Construct) => {
    this._table = new db.Table(scope, "SearchTable", {
      billingMode: db.BillingMode.PAY_PER_REQUEST,
      removalPolicy: RemovalPolicy.DESTROY,
      partitionKey: { name: "PK", type: db.AttributeType.STRING },
      sortKey: { name: "SK", type: db.AttributeType.STRING },
      pointInTimeRecovery: true,
      tableName: "TypsenseDemo",
      stream: db.StreamViewType.NEW_AND_OLD_IMAGES,
    });
  };
}
