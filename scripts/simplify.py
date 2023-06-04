import json

renamed_models = {}


def clean_up_schema_names(schema):
    global renamed_models
    # clean up schemas
    schemas = schema["components"]["schemas"]
    new_schemas = {}
    for name, model in schemas.items():
        new_name = rename_verbose_names(name)
        renamed_models[name] = new_name
        new_schemas[new_name] = model
    schema["components"]["schemas"] = new_schemas


def clean_up_schema_name_references(schema):
    global renamed_models
    if isinstance(schema, dict):
        for p, v in schema.items():
            schema[p] = clean_up_schema_name_references(v)
    elif isinstance(schema, list):
        for i in range(0, len(schema)):
            schema[i] = clean_up_schema_name_references(schema[i])
    elif isinstance(schema, str):
        if schema.startswith("#"):
            ref = schema.split("/")[-1]
            if ref in renamed_models:
                return replace_if_present(schema, ref, renamed_models[ref])
    return schema


def rename_verbose_names(name: str) -> str:
    name = remove_if_present(name, 'ShipBob.')
    name = remove_if_present(name, 'Shipbob.')
    name = remove_if_present(name, 'Microsoft.AspNetCore.Mvc.')
    name = remove_if_present(name, 'Public.Common.Models.')
    name = remove_if_present(name, 'Public.Common.')
    name = remove_if_present(name, 'Public.Api.Models.')
    name = remove_if_present(name, 'Public.Api.ViewModels.')
    name = remove_if_present(name, 'Api.ViewModels.Public.')
    name = remove_if_present(name, 'Api.Models.Public.')
    name = remove_if_present(name, 'Api.ViewModels.')
    name = remove_if_present(name, 'Presentation.Models.')
    name = remove_if_present(name, 'Presentation.ViewModels.')
    name = remove_if_present(name, 'Common.Models.')
    name = remove_if_present(name, 'Domain.Models.')
    name = remove_if_present(name, 'Integrations.')

    # remove common schema suffixes
    name = remove_if_present(name, 'ViewModel')
    name = remove_if_present(name, 'Model')

    # remove "easy" specific schema prefixes
    name = replace_if_present(name, 'Orders.Address', 'OrderAddress')
    name = replace_if_present(name, 'Orders.Inventory', 'OrderInventory')
    name = replace_if_present(name, 'Orders.Estimate', 'OrderEstimate')
    name = replace_if_present(name, 'Orders.Measurements', 'OrderMeasurements')
    name = replace_if_present(name, 'Orders.StatusDetail', 'OrderStatusDetail')
    name = remove_if_present(name, 'Orders.')

    name = replace_if_present(name, 'Location.Address', 'LocationAddress')
    name = replace_if_present(name, 'Location.Region', 'LocationRegion')
    name = replace_if_present(name, 'Location.Service', 'LocationService')
    name = remove_if_present(name, 'Location.')

    name = remove_if_present(name, 'Channels.')

    name = replace_if_present(name, 'Webhooks.CreateWebhookSubscription', 'WebhookSubscription')
    name = remove_if_present(name, 'Webhooks.')

    name = replace_if_present(name, 'Inventory.FulfillmentCenterQuantity',
                              'InventoryFulfillmentCenterQuantity')
    name = replace_if_present(name, 'Inventory.LotQuantity', 'InventoryLotQuantity')
    name = remove_if_present(name, 'Inventory.')

    name = replace_if_present(name, 'Products.Channel', 'ProductChannel')
    name = replace_if_present(name, 'Products.FulfillmentCenterQuantity',
                              'ProductFulfillmentCenterQuantity')
    name = replace_if_present(name, 'Products.InventoryItem', 'ProductInventoryItem')
    name = remove_if_present(name, 'Products.')

    name = replace_if_present(name, 'Receiving.FulfillmentCenter', 'ReceivingFulfillmentCenter')
    name = remove_if_present(name, 'Receiving.')

    name = replace_if_present(name, 'Returns.ChannelInfo', 'ReturnChannelInfo')
    name = replace_if_present(name, 'Returns.FulfillmentCenter', 'ReturnFulfillmentCenter')
    name = replace_if_present(name, 'Returns.InventoryItem', 'ReturnInventoryItem')
    name = replace_if_present(name, 'Returns.Transaction', 'ReturnTransaction')
    name = remove_if_present(name, 'Returns.')

    name = replace_if_present(name, 'Microsoft_AspNetCore_Mvc_ValidationProblemDetails_allOf',
                              'ValidationProblemDetails_allOf')
    return name


def remove_if_present(name: str, word: str) -> str:
    return replace_if_present(name, word, '')


def replace_if_present(name: str, word: str, replace: str) -> str:
    if word in name:
        return name.replace(word, replace)
    return name


def remove_single_oneof_properties(schema):
    if "oneOf" in schema:
        if len(schema["oneOf"]) == 1 and "$ref" in schema["oneOf"][0]:
            schema["$ref"] = schema["oneOf"][0]["$ref"]
            del schema["oneOf"]
    elif "properties" in schema:
        for key, value in schema["properties"].items():
            if "oneOf" in value and len(value["oneOf"]) == 1 and "$ref" in value["oneOf"][0]:
                value["$ref"] = value["oneOf"][0]["$ref"]
                del value["oneOf"]
            else:
                remove_single_oneof_properties(value)


def remove_unneeded_response_schemas(schema):
    if "application/*+json" in schema:
        del schema["application/*+json"]
    if "application/json-patch+json" in schema:
        del schema["application/json-patch+json"]
    if "text/json" in schema:
        del schema["text/json"]
    if "application/pdf" in schema:
        del schema["application/pdf"]


def fix_technical_properties(schema):
    if isinstance(schema, dict):
        if "$type" in schema:
            schema["ttype"] = schema["$type"]
            del schema["$type"]
        if "type" in schema and schema["type"] == "int":
            schema["type"] = "integer"
        if "format" in schema and schema["format"] == "date-time":
            schema["nullable"] = True
        for k, v in schema.items():
            schema[k] = fix_technical_properties(v)
    elif isinstance(schema, list):
        for i in range(0, len(schema)):
            schema[i] = fix_technical_properties(schema[i])
    elif isinstance(schema, str):
        schema = replace_if_present(schema, '$type', 'ttype')
        schema = replace_if_present(schema, '\u201c', '"')
        schema = replace_if_present(schema, '\u201d', '"')
        schema = replace_if_present(schema, '\u2019', '\'')
        schema = replace_if_present(schema, '\u2013', '-')
        schema = replace_if_present(schema, 'â', '"')
        schema = replace_if_present(schema, 'â', '"')
        schema = replace_if_present(schema, 'â', '-')
        schema = replace_if_present(schema, 'â', '\'')
        schema = replace_if_present(schema, '\\r\\n', ' ')
        schema = replace_if_present(schema, '\r\n', ' ')
        schema = replace_if_present(schema, '\n\n', ' ')
        schema = replace_if_present(schema, '\\n\\n', ' ')
        return schema
    return schema


def add_operation_ids(path: str, path_schema: dict):
    """
    Adds a readable operation ID to an endpoint method. This makes the
    generated openapi client methods much cleaner.
    """
    operation_id = ""
    suffix = ""
    for part in path.split("/"):
        if "{" not in part and "1.0" not in part:
            operation_id += part[:1].upper() + part[1:]
        if "2.0" in part:
            suffix = "V2"
    op_map = {
        "get": "get",
        "post": "create",
        "put": "update",
        "patch": "updateFieldsOf",
        "delete": "delete"
    }
    op_id_key = "operationId"
    for method, method_schema in path_schema.items():
        op = op_map.get(method, method)
        if op == "create" and operation_id.endswith("Cancel"):
            op = "cancel"
            operation_id = operation_id[:-len("Cancel")]
        method_schema[op_id_key] = op + operation_id
        if method == "get" and not path.endswith("}"):
            if operation_id.endswith("Statushistory"):
                continue
            if operation_id.endswith("y"):
                method_schema[op_id_key] = method_schema[op_id_key][:-1] + "ies"
            elif not operation_id.endswith("s"):
                method_schema[op_id_key] += "s"
        if suffix:
            method_schema[op_id_key] += suffix


def process_openapi_schema(file_path):
    with open(file_path, "r") as file:
        schema = json.load(file)

        clean_up_schema_names(schema)
        schema = clean_up_schema_name_references(schema)
        schema = fix_technical_properties(schema)

        for n, comp in schema["components"]["schemas"].items():
            remove_single_oneof_properties(comp)
        for path, endpoint in schema["paths"].items():
            add_operation_ids(path, endpoint)
            for m, method in endpoint.items():
                for c, response in method["responses"].items():
                    if "content" in response:
                        remove_unneeded_response_schemas(response["content"])
                        for s in response["content"]:
                            remove_single_oneof_properties(response["content"][s]["schema"])
                if "requestBody" in method:
                    remove_unneeded_response_schemas(method["requestBody"]["content"])
                    for s in method["requestBody"]["content"]:
                        remove_single_oneof_properties(method["requestBody"]["content"][s]["schema"])
        return schema


def save_processed_schema(schema, output_file):
    with open(output_file, "w") as file:
        json.dump(schema, file, indent=2, sort_keys=True)
    print("Schema processed successfully and saved to", output_file)


def main():
    processed_schema = process_openapi_schema("../shipbob_open_api_original.json")
    save_processed_schema(processed_schema, "../shipbob_open_api.json")


if __name__ == "__main__":
    main()
