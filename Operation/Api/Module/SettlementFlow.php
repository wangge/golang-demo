<?php
/**
 * 订单流水
 */

namespace Module 
class SettlementFlow extends ModuleBase
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
        return \Service\Operation\SettlementFlow::instance()->FillSettlement($settlementNo,$paymentChannel,$paymentNo);
    }

    /**
    *统计.
    *
    * @param array $params 筛选条件 .
    * 
    */
    public function statisticList($params)
    {
        return \Service\Operation\SettlementFlow::instance()->statisticList($params);
    }

    /**
    *统计.
    *
    * @param array $params 筛选条件 .
    * 
    */
    public function add($params)
    {
        return \Service\Operation\SettlementFlow::instance()->add($params);
    }

    /**
    *统计.
    *
    * @param array $params 筛选条件 .
    *  @param int id 
    */
    public function update($params,id)
    {
        return \Service\Operation\SettlementFlow::instance()->update($params,id);
    }
}