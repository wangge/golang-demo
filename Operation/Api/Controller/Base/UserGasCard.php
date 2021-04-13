<?php

/**
 * 气表绑定记录
 */

class Controller_Base_UserGasCard extends Controller_Base_Base
{
    

    /**
    * 气表绑定记录.
    */
    public function action_List()
    {
        
        
        $searchParam = [
                "partner_id" => ZKGetPost("partner_id", ""),// 公司id 
                "card_no" => ZKGetPost("card_no", ""),// 缴费号 
                "mobile" => ZKGetPost("mobile", ""),// 手机号 
                "relation" => ZKGetPost("relation", ""),// 户主关系,1-亲戚，2-户主，3-房东，4-朋友，5-其他，11-夫妻，12-父母，13-兄弟姐妹 
                "create_time_start" => ZKGetPost("create_time_start", ""),// 注册时间开始 
                "create_time_end" => ZKGetPost("create_time_end", ""),// 注册时间结算 
        ];
        
        
         $page_no = ZKGetPost("page_no", "1");// 页码 
        
         $page_size = ZKGetPost("page_size", "10");// 每页显示条数 
        
        $res = \Module\UserGasCard\UserGasCard::instance()->List($searchParam,$page_no,$page_size);
        if ($res) {
            $this->responseSuccess($res, '操作成功');
        } else {
            $this->responseFail('操作失败');
        }
    }
}