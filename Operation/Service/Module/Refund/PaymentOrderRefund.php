<?php

namespace Module\Refund; 
/**
 * 退款记录
 */

class PaymentOrderRefund extends \Module\Base
{
    

    /**
    * 退款记录.
    * 
    * @param array $searchParam 过滤条件 .
    * @param int $page_no 页码 .
    * @param int $page_size 每页显示条数 .
    * 
    * @return mixed
    */
    public function List($searchParam,$page_no,$page_size)
    {
        // 初始化变量
        
        
        $partner_id = $searchParam['partner_id'] ?? ''; // 公司id
        $order_sn = $searchParam['order_sn'] ?? ''; // 订单号
        $trade_no = $searchParam['trade_no'] ?? ''; // 支付单号
        $card_no = $searchParam['card_no'] ?? ''; // 缴费号
        $meter_no = $searchParam['meter_no'] ?? ''; // 表具号
        $source = $searchParam['source'] ?? ''; // 表具号
        $payment_method = $searchParam['payment_method'] ?? ''; // 表具号
        $payment_time_start = $searchParam['payment_time_start'] ?? ''; // 支付时间开始
        $payment_time_end = $searchParam['payment_time_end'] ?? ''; // 支付时间结束
        $refund_time_start = $searchParam['refund_time_start'] ?? ''; // 退款发起时间开始
        $refund_time_end = $searchParam['refund_time_end'] ?? ''; // 退款发起时间结束
        $complete_time_start = $searchParam['complete_time_start'] ?? ''; // 退款完成时间开始
        $complete_time_end = $searchParam['complete_time_end'] ?? ''; // 退款完成时间结束
        $outer_card_no = $searchParam['outer_card_no'] ?? ''; // 第三方缴费号
        $nickname = $searchParam['nickname'] ?? ''; // 付款人
        $mobile = $searchParam['mobile'] ?? ''; // 付款人手机号
        $category = $searchParam['category'] ?? ''; // 对接系统1-营收,3-金卡

        // 构建筛选条件
        $cond = [];
        
        
        // 公司id
        if(!empty($partner_id) && $partner_id != ''){
            $cond['partner_id'] = $partner_id;
        }
        // 订单号
        if(!empty($order_sn) && $order_sn != ''){
            $cond['order_sn'] = $order_sn;
        }
        // 支付单号
        if(!empty($trade_no) && $trade_no != ''){
            $cond['trade_no'] = $trade_no;
        }
        // 缴费号
        if(!empty($card_no) && $card_no != ''){
            $cond['card_no'] = $card_no;
        }
        // 表具号
        if(!empty($meter_no) && $meter_no != ''){
            $cond['meter_no'] = $meter_no;
        }
        // 表具号
        if(!empty($source) && $source != '充值渠道，小程序(huixiang_plus_applet)、终端机(arm_terminal)、APP(huixiang_plus_app)'){
            $cond['source'] = $source;
        }
        // 表具号
        if(!empty($payment_method) && $payment_method != '支付方式，用于区分不同渠道不同类型的支付,支付宝(Alipay)、微信(Weixin),四川农商(SCRural)'){
            $cond['payment_method'] = $payment_method;
        }
        // 支付时间开始
        if(!empty($payment_time_start) && $payment_time_start != ''){
            $cond['payment_time_start'] = $payment_time_start;
        }
        // 支付时间结束
        if(!empty($payment_time_end) && $payment_time_end != ''){
            $cond['payment_time_end'] = $payment_time_end;
        }
        // 退款发起时间开始
        if(!empty($refund_time_start) && $refund_time_start != ''){
            $cond['refund_time_start'] = $refund_time_start;
        }
        // 退款发起时间结束
        if(!empty($refund_time_end) && $refund_time_end != ''){
            $cond['refund_time_end'] = $refund_time_end;
        }
        // 退款完成时间开始
        if(!empty($complete_time_start) && $complete_time_start != ''){
            $cond['complete_time_start'] = $complete_time_start;
        }
        // 退款完成时间结束
        if(!empty($complete_time_end) && $complete_time_end != ''){
            $cond['complete_time_end'] = $complete_time_end;
        }
        // 第三方缴费号
        if(!empty($outer_card_no) && $outer_card_no != ''){
            $cond['outer_card_no'] = $outer_card_no;
        }
        // 付款人
        if(!empty($nickname) && $nickname != ''){
            $cond['nickname'] = $nickname;
        }
        // 付款人手机号
        if(!empty($mobile) && $mobile != ''){
            $cond['mobile'] = $mobile;
        }
        // 对接系统1-营收,3-金卡
        if(!empty($category) && $category != ''){
            $cond['category'] = $category;
        }

        // 构建数据
        $now  = \Util\TimeUtil::now();
        
             $data = [
             "partner_id" => $partner_id, // 公司id
             "order_sn" => $order_sn, // 订单号
             "trade_no" => $trade_no, // 支付单号
             "card_no" => $card_no, // 缴费号
             "meter_no" => $meter_no, // 表具号
             "source" => $source, // 表具号
             "payment_method" => $payment_method, // 表具号
             "payment_time_start" => strtotime($payment_time_start), // 支付时间开始
             "payment_time_end" => strtotime($payment_time_end), // 支付时间结束
             "refund_time_start" => strtotime($refund_time_start), // 退款发起时间开始
             "refund_time_end" => strtotime($refund_time_end), // 退款发起时间结束
             "complete_time_start" => strtotime($complete_time_start), // 退款完成时间开始
             "complete_time_end" => strtotime($complete_time_end), // 退款完成时间结束
             "outer_card_no" => $outer_card_no, // 第三方缴费号
             "nickname" => $nickname, // 付款人
             "mobile" => $mobile, // 付款人手机号
             "category" => $category, // 对接系统1-营收,3-金卡
             "create_time" => $now,
             "update_time" => $now
        ];
    }
}