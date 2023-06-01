import unittest

import simplify


class TestAddOperationId(unittest.TestCase):
    def setUp(self):
        self.path_schema = {
            "get": {},
            "put": {},
            "patch": {},
            "post": {},
            "delete": {},
        }

    def test_orders_case(self):
        simplify.add_operation_ids("/order", self.path_schema)
        self.assertEqual(self.path_schema["get"]["operationId"], "getOrders")
        self.assertEqual(self.path_schema["put"]["operationId"], "updateOrder")
        self.assertEqual(self.path_schema["patch"]["operationId"], "updateFieldsOfOrder")
        self.assertEqual(self.path_schema["post"]["operationId"], "createOrder")
        self.assertEqual(self.path_schema["delete"]["operationId"], "deleteOrder")

    def test_order_case(self):
        simplify.add_operation_ids("/order/{id}", self.path_schema)
        self.assertEqual(self.path_schema["get"]["operationId"], "getOrder")
        self.assertEqual(self.path_schema["put"]["operationId"], "updateOrder")
        self.assertEqual(self.path_schema["patch"]["operationId"], "updateFieldsOfOrder")
        self.assertEqual(self.path_schema["post"]["operationId"], "createOrder")
        self.assertEqual(self.path_schema["delete"]["operationId"], "deleteOrder")

    def test_inventories_case(self):
        simplify.add_operation_ids("/inventory", self.path_schema)
        self.assertEqual(self.path_schema["get"]["operationId"], "getInventory")
        self.assertEqual(self.path_schema["put"]["operationId"], "updateInventory")
        self.assertEqual(self.path_schema["patch"]["operationId"], "updateFieldsOfInventory")
        self.assertEqual(self.path_schema["post"]["operationId"], "createInventory")
        self.assertEqual(self.path_schema["delete"]["operationId"], "deleteInventory")

    def test_inventory_case(self):
        simplify.add_operation_ids("/inventory/{id}", self.path_schema)
        self.assertEqual(self.path_schema["get"]["operationId"], "getInventory")
        self.assertEqual(self.path_schema["put"]["operationId"], "updateInventory")
        self.assertEqual(self.path_schema["patch"]["operationId"], "updateFieldsOfInventory")
        self.assertEqual(self.path_schema["post"]["operationId"], "createInventory")
        self.assertEqual(self.path_schema["delete"]["operationId"], "deleteInventory")

    def test_multi_part_path_case(self):
        simplify.add_operation_ids("/inventory/{id}/logs/{logId}/item", self.path_schema)
        self.assertEqual(self.path_schema["get"]["operationId"], "getInventoryLogsItems")
        self.assertEqual(self.path_schema["put"]["operationId"], "updateInventoryLogsItem")
        self.assertEqual(self.path_schema["patch"]["operationId"], "updateFieldsOfInventoryLogsItem")
        self.assertEqual(self.path_schema["post"]["operationId"], "createInventoryLogsItem")
        self.assertEqual(self.path_schema["delete"]["operationId"], "deleteInventoryLogsItem")


class TestRemoveOneOf(unittest.TestCase):
    def test_single_one_of(self):
        schema = {
            "oneOf": [{"$ref": "test"}]
        }
        simplify.remove_single_oneof_properties(schema)
        self.assertEqual(schema, {
            "$ref": "test"
        })

    def test_nested_single_one_of(self):
        schema = {
            "properties": {
                "prop": {"oneOf": [{"$ref": "test"}]}
            }
        }
        simplify.remove_single_oneof_properties(schema)
        self.assertEqual(schema, {
            "properties": {
                "prop": {"$ref": "test"}
            }
        })

    def test_multi_one_of(self):
        schema = {
            "oneOf": [
                {"$ref": "test"},
                {"$ref": "test"}
            ]
        }
        simplify.remove_single_oneof_properties(schema)
        self.assertEqual(schema, {
            "oneOf": [
                {"$ref": "test"},
                {"$ref": "test"}
            ]
        })


class TestCleanUpSchemaNames(unittest.TestCase):
    def test_replace_names(self):
        schema = {
            "components": {
                "schemas": {
                    "Integrations.Location.Public.Api.ViewModels.AddressViewModel": {},
                    "Integrations.Location.Public.Api.ViewModels.LocationViewModel": {},
                    "Integrations.Location.Public.Api.ViewModels.RegionViewModel": {},
                    "Integrations.Location.Public.Api.ViewModels.ServiceViewModel": {},
                    "Integrations.Location.Public.Common.ServiceTypeEnum": {},
                    "Microsoft.AspNetCore.Mvc.ProblemDetails": {},
                    "Microsoft.AspNetCore.Mvc.ValidationProblemDetails": {},
                    "ShipBob.Channels.Api.ViewModels.ChannelViewModel": {},
                    "ShipBob.Orders.Presentation.Models.AddProductToOrderByProductIdModel": {},
                    "ShipBob.Orders.Presentation.Models.AddProductToOrderByReferenceIdModel": {},
                    "ShipBob.Orders.Presentation.Models.AddProductToOrderModel": {},
                    "ShipBob.Orders.Presentation.Models.CancelShipmentsModel": {},
                    "ShipBob.Orders.Presentation.Models.CreateOrderModel": {},
                    "ShipBob.Orders.Presentation.Models.EstimateFulfillmentRequestModel": {},
                    "ShipBob.Orders.Presentation.Models.EstimateProductInfoModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.AddressViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.CanceledOrderViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.CanceledShipmentViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.CanceledShipmentsViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.ChannelInfoViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.EstimateDetailViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.EstimateViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.EstimationAddressViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.FulfillmentCenterViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.InventoryViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.MeasurementsViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.OrderViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.ProductInfoViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.RecipientInfoViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.RecipientViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.ServiceLevelDetailViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.ShipMethodDetailViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.ShipmentLogViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.ShipmentProductViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.ShipmentViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.StatusDetailViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.TagViewModel": {},
                    "ShipBob.Orders.Presentation.ViewModels.TrackingViewModel": {},
                    "ShipBob.Webhooks.Public.Api.Models.CreateWebhookSubscriptionModel": {},
                    "ShipBob.Webhooks.Public.Api.Models.WebhookViewModel": {},
                    "ShipBob.Webhooks.Public.Common.Topics": {},
                    "Shipbob.Inventory.Api.ViewModels.DimensionViewModel": {},
                    "Shipbob.Inventory.Api.ViewModels.FulfillmentCenterQuantityViewModel": {},
                    "Shipbob.Inventory.Api.ViewModels.InventoryViewModel": {},
                    "Shipbob.Inventory.Api.ViewModels.LotQuantityViewModel": {},
                    "Shipbob.Products.Api.Models.Public.CreateProductModel": {},
                    "Shipbob.Products.Api.Models.Public.UpdateProductModel": {},
                    "Shipbob.Products.Api.ViewModels.Public.BundleRootInformationViewModel": {},
                    "Shipbob.Products.Api.ViewModels.Public.ChannelViewModel": {},
                    "Shipbob.Products.Api.ViewModels.Public.FulfillmentCenterQuantityViewModel": {},
                    "Shipbob.Products.Api.ViewModels.Public.InventoryItemViewModel": {},
                    "Shipbob.Products.Api.ViewModels.Public.ProductViewModel": {},
                    "Shipbob.Products.Common.Models.ProductActiveStatus": {},
                    "Shipbob.Products.Common.Models.ProductBundleStatus": {},
                    "Shipbob.Receiving.Public.Api.Models.AddBoxItemToBoxModel": {},
                    "Shipbob.Receiving.Public.Api.Models.AddBoxToOrderModel": {},
                    "Shipbob.Receiving.Public.Api.Models.AssignOrderToFulfillmentCenterModel": {},
                    "Shipbob.Receiving.Public.Api.Models.BoxItemViewModel": {},
                    "Shipbob.Receiving.Public.Api.Models.BoxViewModel": {},
                    "Shipbob.Receiving.Public.Api.Models.CreateReceivingOrderModel": {},
                    "Shipbob.Receiving.Public.Api.Models.FulfillmentCenterViewModel": {},
                    "Shipbob.Receiving.Public.Api.Models.ReceivingOrderViewModel": {},
                    "Shipbob.Receiving.Public.Common.Models.BoxStatus": {},
                    "Shipbob.Receiving.Public.Common.Models.PackageType": {},
                    "Shipbob.Receiving.Public.Common.Models.PackingType": {},
                    "Shipbob.Receiving.Public.Common.Models.ReceivingStatus": {},
                    "Shipbob.Returns.Public.Api.ViewModels.ChannelInfoViewModel": {},
                    "Shipbob.Returns.Public.Api.ViewModels.CreateReturnViewModel": {},
                    "Shipbob.Returns.Public.Api.ViewModels.FulfillmentCenterViewModel": {},
                    "Shipbob.Returns.Public.Api.ViewModels.InventoryItemViewModel": {},
                    "Shipbob.Returns.Public.Api.ViewModels.ReturnActionRequestedViewModel": {},
                    "Shipbob.Returns.Public.Api.ViewModels.ReturnActionTakenViewModel": {},
                    "Shipbob.Returns.Public.Api.ViewModels.ReturnInventoryViewModel": {},
                    "Shipbob.Returns.Public.Api.ViewModels.ReturnOrderStatusHistoryViewModel": {},
                    "Shipbob.Returns.Public.Api.ViewModels.ReturnOrderViewModel": {},
                    "Shipbob.Returns.Public.Api.ViewModels.TransactionViewModel": {},
                    "Shipbob.Returns.Public.Common.ReturnAction": {},
                    "Shipbob.Returns.Public.Common.ReturnActionSource": {},
                    "Shipbob.Returns.Public.Common.ReturnStatus": {},
                    "Shipbob.Returns.Public.Common.ReturnType": {},
                    "Shipbob.Returns.Public.Common.SortOrder": {},
                    "Shipbob.Returns.Public.Common.TransactionLogSource": {},
                }
            }
        }
        simplify.clean_up_schema_names(schema)
        self.assertEqual({
            "components": {
                "schemas": {
                    "LocationAddress": {},
                    "Location": {},
                    "LocationRegion": {},
                    "LocationService": {},
                    "LocationServiceTypeEnum": {},
                    "ProblemDetails": {},
                    "ValidationProblemDetails": {},
                    "Channel": {},
                    "AddProductToOrderByProductIdModel": {},
                    "AddProductToOrderByReferenceIdModel": {},
                    "AddProductToOrderModel": {},
                    "CancelShipmentsModel": {},
                    "CreateOrderModel": {},
                    "OrderEstimateFulfillmentRequestModel": {},
                    "OrderEstimateProductInfoModel": {},
                    "OrderAddress": {},
                    "CanceledOrder": {},
                    "CanceledShipment": {},
                    "CanceledShipments": {},
                    "ChannelInfo": {},
                    "OrderEstimateDetail": {},
                    "OrderEstimate": {},
                    "EstimationAddress": {},
                    "FulfillmentCenter": {},
                    "OrderInventory": {},
                    "OrderMeasurements": {},
                    "Order": {},
                    "ProductInfo": {},
                    "RecipientInfo": {},
                    "Recipient": {},
                    "ServiceLevelDetail": {},
                    "ShipMethodDetail": {},
                    "ShipmentLog": {},
                    "ShipmentProduct": {},
                    "Shipment": {},
                    "OrderStatusDetail": {},
                    "Tag": {},
                    "Tracking": {},
                    "WebhookSubscriptionModel": {},
                    "Webhook": {},
                    "Topics": {},
                    "Dimension": {},
                    "InventoryFulfillmentCenterQuantity": {},
                    "Inventory": {},
                    "InventoryLotQuantity": {},
                    "CreateProductModel": {},
                    "UpdateProductModel": {},
                    "BundleRootInformation": {},
                    "ProductChannel": {},
                    "ProductFulfillmentCenterQuantity": {},
                    "ProductInventoryItem": {},
                    "Product": {},
                    "ProductActiveStatus": {},
                    "ProductBundleStatus": {},
                    "AddBoxItemToBoxModel": {},
                    "AddBoxToOrderModel": {},
                    "AssignOrderToFulfillmentCenterModel": {},
                    "BoxItem": {},
                    "Box": {},
                    "CreateReceivingOrderModel": {},
                    "ReceivingFulfillmentCenter": {},
                    "ReceivingOrder": {},
                    "BoxStatus": {},
                    "PackageType": {},
                    "PackingType": {},
                    "ReceivingStatus": {},
                    "ReturnChannelInfo": {},
                    "CreateReturn": {},
                    "ReturnFulfillmentCenter": {},
                    "ReturnInventoryItem": {},
                    "ReturnActionRequested": {},
                    "ReturnActionTaken": {},
                    "ReturnInventory": {},
                    "ReturnOrderStatusHistory": {},
                    "ReturnOrder": {},
                    "ReturnTransaction": {},
                    "ReturnAction": {},
                    "ReturnActionSource": {},
                    "ReturnStatus": {},
                    "ReturnType": {},
                    "SortOrder": {},
                    "ReturnTransactionLogSource": {},
                }
            }
        }, schema)


if __name__ == '__main__':
    unittest.main()
