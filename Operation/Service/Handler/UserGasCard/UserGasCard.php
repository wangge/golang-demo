<?php

namespace Handler\UserGasCard; 
/**
 * 气表绑定记录
 */

class UserGasCard extends \Handler\Base
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
        return \Module\UserGasCard\UserGasCard::instance()->List($searchParam,$page_no,$page_size);
    }
}