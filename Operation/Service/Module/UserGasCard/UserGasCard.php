<?php

namespace Module\UserGasCard; 
/**
 * 气表绑定记录
 */

class UserGasCard extends \Module\Base
{
    

    /**
    * 气表绑定记录.
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
        $card_no = $searchParam['card_no'] ?? ''; // 缴费号
        $mobile = $searchParam['mobile'] ?? ''; // 手机号
        $relation = $searchParam['relation'] ?? ''; // 户主关系,1-亲戚，2-户主，3-房东，4-朋友，5-其他，11-夫妻，12-父母，13-兄弟姐妹
        $create_time_start = $searchParam['create_time_start'] ?? ''; // 注册时间开始
        $create_time_end = $searchParam['create_time_end'] ?? ''; // 注册时间结算

        // 构建筛选条件
        $cond = [];
        
        
        // 公司id
        if(!empty($partner_id) && $partner_id != ''){
            $cond['partner_id'] = $partner_id;
        }
        // 缴费号
        if(!empty($card_no) && $card_no != ''){
            $cond['card_no'] = $card_no;
        }
        // 手机号
        if(!empty($mobile) && $mobile != ''){
            $cond['mobile'] = $mobile;
        }
        // 户主关系,1-亲戚，2-户主，3-房东，4-朋友，5-其他，11-夫妻，12-父母，13-兄弟姐妹
        if(!empty($relation) && $relation != ''){
            $cond['relation'] = $relation;
        }
        // 注册时间开始
        if(!empty($create_time_start) && $create_time_start != ''){
            $cond['create_time_start'] = $create_time_start;
        }
        // 注册时间结算
        if(!empty($create_time_end) && $create_time_end != ''){
            $cond['create_time_end'] = $create_time_end;
        }

        // 构建数据
        $now  = \Util\TimeUtil::now();
        
             $data = [
             "partner_id" => $partner_id, // 公司id
             "card_no" => $card_no, // 缴费号
             "mobile" => $mobile, // 手机号
             "relation" => $relation, // 户主关系,1-亲戚，2-户主，3-房东，4-朋友，5-其他，11-夫妻，12-父母，13-兄弟姐妹
             "create_time_start" => strtotime($create_time_start), // 注册时间开始
             "create_time_end" => strtotime($create_time_end), // 注册时间结算
             "create_time" => $now,
             "update_time" => $now
        ];
    }
}