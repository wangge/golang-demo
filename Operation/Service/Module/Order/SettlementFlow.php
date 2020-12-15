<?php
/**
 * 订单流水
 */

namespace Module\Order 
class SettlementFlow extends \Module\Base
{
    

    /**
    *完成划账.
    *
    * @param string $settlementNo 结算单号 .
    * @param string $paymentChannel 渠道 .
    * @param string $paymentNo 流水号 .
    * 
    */
    public function FillSettlement($settlementNo,$paymentChannel,$paymentNo)
    {
        
    }

    /**
    *统计.
    *
    * @param array $params 筛选条件 .
    * 
    */
    public function statisticList($params)
    {
        
        $params = [
                "page_no" => ZKGetPost("page_no", "1"),// 页码 
                "page_size" => ZKGetPost("page_size", "10"),// 每页显示条数 
                "payment_time_start" => ZKGetPost("payment_time_start", ""),// 支付开始时间 
        ];
    }

    /**
    *统计.
    *
    * @param array $params 筛选条件 .
    * 
    */
    public function add($params)
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
    }

    /**
    *统计.
    *
    * @param array $params 筛选条件 .
    *  @param int id 
    */
    public function update($params,id)
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
    }
}