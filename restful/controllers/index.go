package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"goshop/restful/models"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (crtl *IndexController) Index(c *gin.Context) {
	systemGroup := &models.SystemGroup{}
	list, err := systemGroup.All()
	if err != nil {
		handleErr(c, err)
		return
	}

	mList := make(map[string]*models.SystemGroup, 0)
	for _, item := range list {

		mList[item.ConfigName] = item
	}

	systemConfig := &models.SystemConfig{}

	mMenuNameValue, err := systemConfig.AllMenuNameKey()
	if err != nil {
		handleErr(c, err)
		return
	}

	sGroupIds := make([]int, 0)
	routineHomeBannerId := mList["routine_home_banner"].Id
	routineHomeMenusId := mList["routine_home_menus"].Id
	routineHomeRollNewsId := mList["routine_home_roll_news"].Id
	routineHomeActivityId := mList["routine_home_activity"].Id
	routineHomeBastBannerId := mList["routine_home_bast_banner"].Id
	sGroupIds = append(sGroupIds, routineHomeBannerId)
	sGroupIds = append(sGroupIds, routineHomeMenusId)
	sGroupIds = append(sGroupIds, routineHomeRollNewsId)
	sGroupIds = append(sGroupIds, routineHomeActivityId)
	sGroupIds = append(sGroupIds, routineHomeBastBannerId)

	routineHomeBannerList := make([]interface{}, 0)
	routineHomeMenusList := make([]interface{}, 0)
	routineHomeRollNewsList := make([]interface{}, 0)
	routineHomeActivityList := make([]interface{}, 0)
	routineHomeBastBannerList := make([]interface{}, 0)

	systemGroupData := &models.SystemGroupData{}

	sGroupDataList, err := systemGroupData.ListByGids(sGroupIds)

	if err != nil {
		handleErr(c, err)
		return
	}

	for _, item := range sGroupDataList {
		if item.Gid == routineHomeBannerId {
			itemdata := make(map[string]interface{}, 0)
			itemdata["id"] = item.Id
			valMap := make(map[string]map[string]string, 0)
			err := json.Unmarshal([]byte(item.Value), &valMap)
			if err != nil {
				handleErr(c, err)
				return
			}
			name := valMap["name"]
			itemdata["name"] = name["value"]

			pic := valMap["pic"]
			itemdata["pic"] = pic["value"]

			urlVal := valMap["url"]
			itemdata["url"] = urlVal["value"]

			routineHomeBannerList = append(routineHomeBannerList, itemdata)
		}
		if item.Gid == routineHomeMenusId {
			itemdata := make(map[string]interface{}, 0)
			itemdata["id"] = item.Id
			valMap := make(map[string]map[string]string, 0)
			err := json.Unmarshal([]byte(item.Value), &valMap)
			if err != nil {
				handleErr(c, err)
				return
			}

			name := valMap["name"]
			itemdata["name"] = name["value"]

			show := valMap["show"]
			itemdata["show"] = show["value"]

			urlVal := valMap["url"]
			itemdata["url"] = urlVal["value"]

			pic := valMap["pic"]
			itemdata["pic"] = pic["value"]

			routineHomeMenusList = append(routineHomeMenusList, itemdata)
		}
		if item.Gid == routineHomeRollNewsId {
			itemdata := make(map[string]interface{}, 0)
			itemdata["id"] = item.Id
			valMap := make(map[string]map[string]string, 0)
			err := json.Unmarshal([]byte(item.Value), &valMap)
			if err != nil {
				handleErr(c, err)
				return
			}

			info := valMap["info"]
			itemdata["info"] = info["value"]

			show := valMap["show"]
			itemdata["show"] = show["value"]

			urlVal := valMap["url"]
			itemdata["url"] = urlVal["value"]

			routineHomeRollNewsList = append(routineHomeRollNewsList, itemdata)
		}
		if item.Gid == routineHomeActivityId {

			itemdata := make(map[string]interface{}, 0)
			itemdata["id"] = item.Id
			valMap := make(map[string]map[string]string, 0)
			err := json.Unmarshal([]byte(item.Value), &valMap)
			if err != nil {
				handleErr(c, err)
				return
			}

			info := valMap["info"]
			itemdata["info"] = info["value"]

			link := valMap["link"]
			itemdata["link"] = link["value"]

			pic := valMap["pic"]
			itemdata["pic"] = pic["value"]

			title := valMap["title"]
			itemdata["title"] = title["value"]

			routineHomeActivityList = append(routineHomeActivityList, itemdata)
		}

		if item.Gid == routineHomeBastBannerId {
			itemdata := make(map[string]interface{}, 0)
			itemdata["id"] = item.Id
			valMap := make(map[string]map[string]string, 0)
			err := json.Unmarshal([]byte(item.Value), &valMap)
			if err != nil {
				handleErr(c, err)
				return
			}

			comment := valMap["comment"]
			itemdata["comment"] = comment["value"]

			link := valMap["link"]
			itemdata["link"] = link["value"]

			img := valMap["img"]
			itemdata["img"] = img["value"]
			routineHomeBastBannerList = append(routineHomeBastBannerList, itemdata)

		}

	}

	mdata := make(map[string]interface{}, 0)

	mdata["banner"] = routineHomeBannerList     //首页banner图
	mdata["menus"] = routineHomeMenusList       //首页按钮
	mdata["roll"] = routineHomeRollNewsList     //首页滚动新闻
	mdata["activity"] = routineHomeActivityList //首页活动区域图片

	info := make(map[string]interface{}, 0)

	var fastInfo interface{}

	err = json.Unmarshal([]byte(mMenuNameValue["fast_info"].Value), &fastInfo)
	if err != nil {
		handleErr(c, err)
		return
	}

	var bastInfo interface{}

	err = json.Unmarshal([]byte(mMenuNameValue["bast_info"].Value), &bastInfo)
	if err != nil {
		handleErr(c, err)
		return
	}
	var firstInfo interface{}

	err = json.Unmarshal([]byte(mMenuNameValue["first_info"].Value), &firstInfo)
	if err != nil {
		handleErr(c, err)
		return
	}
	var salesInfo interface{}

	err = json.Unmarshal([]byte(mMenuNameValue["sales_info"].Value), &salesInfo)
	if err != nil {
		handleErr(c, err)
		return
	}

	var routineIndexLogo interface{}

	err = json.Unmarshal([]byte(mMenuNameValue["routine_index_logo"].Value), &routineIndexLogo)
	if err != nil {
		handleErr(c, err)
		return
	}

	info["fastInfo"] = fastInfo
	info["bastInfo"] = bastInfo
	info["firstInfo"] = firstInfo
	info["salesInfo"] = salesInfo
	info["bastBanner"] = routineHomeBastBannerList

	fastNumberStr := mMenuNameValue["fast_number"].Value
	bastNumberStr := mMenuNameValue["bast_number"].Value
	firstNumberStr := mMenuNameValue["first_number"].Value

	err = json.Unmarshal([]byte(fastNumberStr), &fastNumberStr)
	if err != nil {
		handleErr(c, err)
		return
	}

	err = json.Unmarshal([]byte(bastNumberStr), &bastNumberStr)
	if err != nil {
		handleErr(c, err)
		return
	}

	err = json.Unmarshal([]byte(firstNumberStr), &firstNumberStr)
	if err != nil {
		handleErr(c, err)
		return
	}

	fastNumber, err := strconv.Atoi(fastNumberStr)
	if err != nil {
		handleErr(c, err)
		return
	}

	bastNumber, err := strconv.Atoi(bastNumberStr)
	if err != nil {
		handleErr(c, err)
		return
	}
	firstNumber, err := strconv.Atoi(firstNumberStr)
	if err != nil {
		handleErr(c, err)
		return
	}

	fmt.Println("fastNumberStr: ", fastNumberStr, "fastNumber: ", fastNumber)
	fmt.Println("bastNumberStr: ", bastNumberStr, "bastNumber: ", bastNumber)
	fmt.Println("firstNumberStr: ", firstNumberStr, "firstNumber: ", firstNumber)

	mdata["logoUrl"] = routineIndexLogo

	storeCategory := &models.StoreCategory{}

	storeCategoryList, err := storeCategory.ListByIndex(fastNumber)
	if err != nil {
		handleErr(c, err)
		return
	}
	storeCategoryListMaps := make([]map[string]interface{}, 0)
	for _, item := range storeCategoryList {
		itemdata := make(map[string]interface{}, 0)
		itemdata["cate_name"] = item.CateName
		itemdata["id"] = item.Id
		itemdata["pid"] = item.Pid
		itemdata["pic"] = item.Pic
		storeCategoryListMaps = append(storeCategoryListMaps, itemdata)
	}

	info["fastList"] = storeCategoryListMaps

	storeProduct := &models.StoreProduct{}
	bestProductList, err := storeProduct.GetBestProduct(bastNumber, 0)
	if err != nil {
		handleErr(c, err)
		return
	}

	info["bastList"] = bestProductList

	firstLists, err := storeProduct.GetNewProduct(firstNumber)
	if err != nil {
		handleErr(c, err)
		return
	}
	info["firstList"] = firstLists

	benefitList, err := storeProduct.GetBenefitProduct(3)
	if err != nil {
		handleErr(c, err)
		return
	}

	mdata["benefit"] = benefitList

	likeInfoList, err := storeProduct.GetHotProduct(3, 0)
	if err != nil {
		handleErr(c, err)
		return
	}
	mdata["likeInfo"] = likeInfoList

	mdata["info"] = info

	// @todo  $couponList=StoreCouponIssue::getIssueCouponList($this->uid,3);
	// 优惠券

	handleOk(c, mdata)
}

func (crtl *IndexController) MyNaviga(c *gin.Context) {
	systemGroup := &models.SystemGroup{}
	list, err := systemGroup.All()
	if err != nil {
		handleErr(c, err)
		return
	}

	mList := make(map[string]*models.SystemGroup, 0)
	for _, item := range list {

		mList[item.ConfigName] = item
	}

	sGroupIds := make([]int, 0)
	routineMyMenusId := mList["routine_my_menus"].Id
	sGroupIds = append(sGroupIds, routineMyMenusId)
	systemGroupData := &models.SystemGroupData{}

	routineMyMenusList := make([]interface{}, 0)

	sGroupDataList, err := systemGroupData.ListByGids(sGroupIds)

	if err != nil {
		handleErr(c, err)
		return
	}

	for _, item := range sGroupDataList {
		if item.Gid == routineMyMenusId {

			itemdata := make(map[string]interface{}, 0)
			itemdata["id"] = item.Id
			valMap := make(map[string]map[string]string, 0)
			err := json.Unmarshal([]byte(item.Value), &valMap)
			if err != nil {
				handleErr(c, err)
				return
			}

			name := valMap["name"]
			itemdata["name"] = name["value"]

			pic := valMap["pic"]
			itemdata["pic"] = pic["value"]

			urlVal := valMap["url"]
			itemdata["url"] = urlVal["value"]
			routineMyMenusList = append(routineMyMenusList, itemdata)
		}
	}

	mdata := make(map[string]interface{}, 0)
	mdata["routine_my_menus"] = routineMyMenusList
	handleOk(c, mdata)

}

//获取首页推荐不同类型产品的轮播图和产品
func (crtl *IndexController) GetIndexGroomList(c *gin.Context) {
	typStr := c.Param("typ")

	typ, err := strconv.Atoi(typStr)
	if err != nil {
		handleErr(c, err)
		return
	}

	name := "routine_home_bast_banner"

	if typ == 2 {
		name = "routine_home_hot_banner"
	}
	if typ == 3 {
		name = "routine_home_new_banner"
	}
	if typ == 4 {
		name = "routine_home_benefit_banner"
	}

	systemGroup := &models.SystemGroup{}

	err = systemGroup.GetByConfigName(name)
	if err != nil {
		handleErr(c, err)
		return
	}

	ids := []int{systemGroup.Id}

	systemGroupData := &models.SystemGroupData{}

	sGroupDataList, err := systemGroupData.ListByGids(ids)

	if err != nil {
		handleErr(c, err)
		return
	}

	branners := make([]interface{}, 0)
	for _, item := range sGroupDataList {
		itemdata := make(map[string]interface{}, 0)
		itemdata["id"] = item.Id
		valMap := make(map[string]map[string]string, 0)
		err := json.Unmarshal([]byte(item.Value), &valMap)
		if err != nil {
			handleErr(c, err)
			return
		}

		comment := valMap["comment"]
		itemdata["comment"] = comment["value"]

		img := valMap["img"]
		itemdata["img"] = img["value"]

		branners = append(branners, itemdata)
	}

	mdata := make(map[string]interface{}, 0)

	mdata["branner"] = branners

	storeProduct := &models.StoreProduct{}

	if typ == 1 {
		bestProductList, err := storeProduct.GetBestProduct(100, 0)
		if err != nil {
			handleErr(c, err)
			return
		}

		mdata["list"] = bestProductList
	}

	if typ == 2 {
		likeInfoList, err := storeProduct.GetHotProduct(100, 0)
		if err != nil {
			handleErr(c, err)
			return
		}
		mdata["list"] = likeInfoList
	}

	if typ == 3 {
		firstLists, err := storeProduct.GetNewProduct(100)
		if err != nil {
			handleErr(c, err)
			return
		}
		mdata["list"] = firstLists
	}

	if typ == 4 {
		benefitList, err := storeProduct.GetBenefitProduct(100)
		if err != nil {
			handleErr(c, err)
			return
		}
		mdata["list"] = benefitList
	}
	handleOk(c, mdata)

}

func (crtl *IndexController) SystemGroupDataValue(c *gin.Context) {

	name := c.Param("name")
	systemGroup := &models.SystemGroup{}
	err := systemGroup.GetByConfigName(name)
	if err != nil {
		handleErr(c, err)
		return
	}

	ids := []int{systemGroup.Id}

	systemGroupData := &models.SystemGroupData{}

	sGroupDataList, err := systemGroupData.ListByGids(ids)

	if err != nil {
		handleErr(c, err)
		return
	}
	mdata := make(map[string]interface{}, 0)

	if name == "sign_day_num" {
		sdayNums := make([]interface{}, 0)
		for _, item := range sGroupDataList {
			itemdata := make(map[string]interface{}, 0)
			itemdata["id"] = item.Id
			valMap := make(map[string]map[string]string, 0)
			err := json.Unmarshal([]byte(item.Value), &valMap)
			if err != nil {
				handleErr(c, err)
				return
			}

			day := valMap["day"]
			itemdata["day"] = day["value"]

			sign_num := valMap["sign_num"]
			itemdata["sign_num"] = sign_num["value"]

			sdayNums = append(sdayNums, itemdata)
		}
		mdata["sign_day_num"] = sdayNums
	}

	handleOk(c, mdata)

}
