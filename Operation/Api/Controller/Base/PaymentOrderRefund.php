<?php

/**
 * 退款记录
 */

class Controller_Base_PaymentOrderRefund extends Controller_Base_Base
{
    

    /**
    * 退款记录.
    */
    public function action_List()
    {
        
        
        $searchParam = [
                "partner_id" => ZKGetPost("partner_id", ""),// 公司id 
                "order_sn" => ZKGetPost("order_sn", ""),// 订单号 
                "trade_no" => ZKGetPost("trade_no", ""),// 支付单号 
                "card_no" => ZKGetPost("card_no", ""),// 缴费号 
                "meter_no" => ZKGetPost("meter_no", ""),// 表具号 
                "source" => ZKGetPost("source", "充值渠道，小程序(huixiang_plus_applet)、终端机(arm_terminal)、APP(huixiang_plus_app)"),// 表具号 
                "payment_method" => ZKGetPost("payment_method", "支付方式，用于区分不同渠道不同类型的支付,支付宝(Alipay)、微信(Weixin),四川农商(SCRural)"),// 表具号 
                "payment_time_start" => ZKGetPost("payment_time_start", ""),// 支付时间开始 
                "payment_time_end" => ZKGetPost("payment_time_end", ""),// 支付时间结束 
                "refund_time_start" => ZKGetPost("refund_time_start", ""),// 退款发起时间开始 
                "refund_time_end" => ZKGetPost("refund_time_end", ""),// 退款发起时间结束 
                "complete_time_start" => ZKGetPost("complete_time_start", ""),// 退款完成时间开始 
                "complete_time_end" => ZKGetPost("complete_time_end", ""),// 退款完成时间结束 
                "outer_card_no" => ZKGetPost("outer_card_no", ""),// 第三方缴费号 
                "nickname" => ZKGetPost("nickname", ""),// 付款人 
                "mobile" => ZKGetPost("mobile", ""),// 付款人手机号 
                "category" => ZKGetPost("category", ""),// 对接系统1-营收,3-金卡 
        ];
        
        
         $page_no = ZKGetPost("page_no", "1");// 页码 
        
         $page_size = ZKGetPost("page_size", "10");// 每页显示条数 
        
        $res = \Module\Refund\PaymentOrderRefund::instance()->List($searchParam,$page_no,$page_size);
        if ($res) {
            $this->responseSuccess($res, '操作成功');
        } else {
            $this->responseFail('操作失败');
        }
    }
}