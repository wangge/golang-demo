<?php
/**
 * 订单流水
 */

class Controller_Base_SettlementFlow extends Controller_Base_Base
{
    

    /**
    *完成划账.
    */
    public function action_FillSettlement()
    {
        
         $settlementNo = ZKGetPost("settlementNo", "");// 结算单号 
        
         $paymentChannel = ZKGetPost("paymentChannel", "");// 渠道 
        
         $paymentNo = ZKGetPost("paymentNo", "");// 流水号 
        
        $res = \Module\SettlementFlow::instance()->FillSettlement($settlementNo,$paymentChannel,$paymentNo);
        if ($res) {
            $this->responseSuccess($res, '操作成功');
        } else {
            $this->responseFail('操作失败');
        }
    }

    /**
    *统计.
    */
    public function action_statisticList()
    {
        
        
        $params = [
                "page_no" => ZKGetPost("page_no", "1"),// 页码 
                "page_size" => ZKGetPost("page_size", "10"),// 每页显示条数 
                "payment_time_start" => ZKGetPost("payment_time_start", ""),// 支付开始时间 
        ];
        
        
        $res = \Module\SettlementFlow::instance()->statisticList($params);
        if ($res) {
            $this->responseSuccess($res, '操作成功');
        } else {
            $this->responseFail('操作失败');
        }
    }

    /**
    *统计.
    */
    public function action_add()
    {
        
        
        $params = [
                "id" => ZKGetPost("id", ""),// 自增ID 
                "settlement_no" => ZKGetPost("settlement_no", ""),// 结算单号 
                "partner_id" => ZKGetPost("partner_id", ""),// 商家ID 
                "store_id" => ZKGetPost("store_id", ""),// 店铺ID 
                "order_num" => ZKGetPost("order_num", ""),// (付款)订单数 
                "refund_num" => ZKGetPost("refund_num", ""),// (退货/退款)订单数 
                "order_amount" => ZKGetPost("order_amount", ""),// (付款)订单总金额 
                "refund_amount" => ZKGetPost("refund_amount", ""),// (退货/退款)总金额 
                "settlement_amount" => ZKGetPost("settlement_amount", ""),// 结算金额 
                "settlement_rate" => ZKGetPost("settlement_rate", ""),// 结算费率，表示千分之几 
                "settlement_fee" => ZKGetPost("settlement_fee", ""),// 结算手续费 
                "operator" => ZKGetPost("operator", ""),// 结算单创建人 
                "status" => ZKGetPost("status", ""),// 结算单状态：1-未划账（新创建）， 2-已划账 
                "payment_bill_id" => ZKGetPost("payment_bill_id", ""),// 付款单ID 
                "create_time" => ZKGetPost("create_time", ""),// 创建时间 
                "update_time" => ZKGetPost("update_time", ""),// 最后更新时间 
        ];
        
        
        $res = \Module\SettlementFlow::instance()->add($params);
        if ($res) {
            $this->responseSuccess($res, '操作成功');
        } else {
            $this->responseFail('操作失败');
        }
    }

    /**
    *统计.
    */
    public function action_update()
    {
        
        
        $params = [
                "id" => ZKGetPost("id", ""),// 自增ID 
                "settlement_no" => ZKGetPost("settlement_no", ""),// 结算单号 
                "partner_id" => ZKGetPost("partner_id", ""),// 商家ID 
                "store_id" => ZKGetPost("store_id", ""),// 店铺ID 
                "order_num" => ZKGetPost("order_num", ""),// (付款)订单数 
                "refund_num" => ZKGetPost("refund_num", ""),// (退货/退款)订单数 
                "order_amount" => ZKGetPost("order_amount", ""),// (付款)订单总金额 
                "refund_amount" => ZKGetPost("refund_amount", ""),// (退货/退款)总金额 
                "settlement_amount" => ZKGetPost("settlement_amount", ""),// 结算金额 
                "settlement_rate" => ZKGetPost("settlement_rate", ""),// 结算费率，表示千分之几 
                "settlement_fee" => ZKGetPost("settlement_fee", ""),// 结算手续费 
                "operator" => ZKGetPost("operator", ""),// 结算单创建人 
                "status" => ZKGetPost("status", ""),// 结算单状态：1-未划账（新创建）， 2-已划账 
                "payment_bill_id" => ZKGetPost("payment_bill_id", ""),// 付款单ID 
                "create_time" => ZKGetPost("create_time", ""),// 创建时间 
                "update_time" => ZKGetPost("update_time", ""),// 最后更新时间 
        ];
        
        
        $res = \Module\SettlementFlow::instance()->update($params,id);
        if ($res) {
            $this->responseSuccess($res, '操作成功');
        } else {
            $this->responseFail('操作失败');
        }
    }
}