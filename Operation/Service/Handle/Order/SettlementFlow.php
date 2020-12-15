<?php
/**
 * 订单流水
 */

namespace Handle\Order 
class SettlementFlow extends \Handle\Base
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
        return \Module\Order\SettlementFlow::instance()->FillSettlement($settlementNo,$paymentChannel,$paymentNo);
    }

    /**
    *统计.
    *
    * @param array $params 筛选条件 .
    * 
    */
    public function statisticList($params)
    {
        return \Module\Order\SettlementFlow::instance()->statisticList($params);
    }

    /**
    *统计.
    *
    * @param array $params 筛选条件 .
    * 
    */
    public function add($params)
    {
        return \Module\Order\SettlementFlow::instance()->add($params);
    }

    /**
    *统计.
    *
    * @param array $params 筛选条件 .
    *  @param int id 
    */
    public function update($params,id)
    {
        return \Module\Order\SettlementFlow::instance()->update($params,id);
    }
}