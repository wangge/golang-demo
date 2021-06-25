<?php

namespace Module\Refund; 
/**
 * 退款记录
 */

class PaymentOrderRefund extends \Module\ModuleBase
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
        return \Service\Operation\PaymentOrderRefund::instance()->List($searchParam,$page_no,$page_size);
    }
}