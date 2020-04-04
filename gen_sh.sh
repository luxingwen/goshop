#!/bin/sh

rm -rf restful/controllers/gen_*
rm -rf restful/models/gen_*
rm -rf restful/routers/gen_router.go

cd tools/gen_model/
go run main.go
cd ../../
cp tools/gen_model/gen/*.go restful/models/
cp tools/gen_model/gen_controllers/*.go restful/controllers/
cp tools/gen_model/gen_router/*.go restful/routers/
rm -rf restful/controllers/gen_store_combination_attr.go
rm -rf restful/controllers/gen_store_combination_attr_result.go
rm -rf restful/controllers/gen_store_combination_attr_value.go
rm -rf restful/controllers/gen_store_coupon_issue_user.go
rm -rf restful/controllers/gen_store_order_cart_info.go
rm -rf restful/controllers/gen_store_order_status.go
rm -rf restful/controllers/gen_store_product_attr.go
rm -rf restful/controllers/gen_store_product_attr_result.go
rm -rf restful/controllers/gen_store_product_attr_value.go
rm -rf restful/controllers/gen_store_product_relation.go
rm -rf restful/controllers/gen_store_seckill_attr.go
rm -rf restful/controllers/gen_store_seckill_attr_result.go
rm -rf restful/controllers/gen_store_seckill_attr_value.go
rm -rf restful/controllers/gen_system_attachment.go
rm -rf restful/controllers/gen_wechat_user.go
rm -rf restful/controllers/gen_article_content.go
rm -rf restful/controllers/gen_cache.go
rm -rf restful/controllers/gen_user.go