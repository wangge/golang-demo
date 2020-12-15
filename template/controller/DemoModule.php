<?php


namespace Module;


class {NAME} extends ModuleBase
{
    /**
     * 创建{NAME}.
     *
     * @param int   $opUid  操作用户UID.
     * @param array $params {NAME}参数.
     *
     * @return mixed
     */
    public function create{NAME}(int $opUid, array $params)
    {
        return \Service\Operation\{NAME}::instance()->create{NAME}($opUid, $params);
    }

    /**
     * {NAME}详情.
     *
     * @param int ${NAME-LOW}Id 活动ID.
     *
     * @return mixed
     */
    public function get{NAME}Detail(int ${NAME-LOW}Id)
    {
        return \Service\Operation\{NAME}::instance()->get{NAME}Detail(${NAME-LOW}Id);
    }

    /**
     * {NAME}列表.
     *
     * @param array $searchParams 搜索条件.
     * @param int   $pageNo       分页页码.
     * @param int   $pageSize     分页大小.
     *
     * @return mixed
     */
    public function get{NAME}List(array $searchParams = [], $pageNo = 1, $pageSize = 10)
    {
        return \Service\Operation\{NAME}::instance()->get{NAME}List($searchParams, $pageNo, $pageSize);
    }

    /**
     * 删除{NAME}.
     *
     * @param int $opUid      操作用户UID.
     * @param int ${NAME-LOW}Id 活动ID.
     *
     * @return mixed
     */
    public function delete{NAME}(int $opUid, int ${NAME-LOW}Id)
    {
        return \Service\Operation\{NAME}::instance()->delete{NAME}($opUid, ${NAME-LOW}Id);
    }

    /**
     * 编辑{NAME}.
     *
     * @param int   $opUid      操作用户UID.
     * @param int   ${NAME-LOW}Id {NAME-LOW}Id.
     * @param array $params     {NAME}参数.
     *
     * @return mixed
     */
    public function edit{NAME}(int $opUid, int ${NAME-LOW}Id, array $params)
    {
        $params['start_time'] = strtotime($params['start_time']);
        $params['end_time'] = strtotime($params['end_time']);
        $params['prize'] = json_decode($params['prize'], true) ?: [];
        return \Service\Operation\{NAME}::instance()->edit{NAME}($opUid, ${NAME-LOW}Id, $params);
    }

    /**
     * 上线{NAME}.
     *
     * @param int $opUid      商户后台用户UID.
     * @param int ${NAME-LOW}Id {NAME-LOW}Id.
     *
     * @return mixed
     */
    public function doOnline($opUid, ${NAME-LOW}Id)
    {
        return \Service\Operation\{NAME}::instance()->doOnline($opUid, ${NAME-LOW}Id);
    }

    /**
     * 下线/结束{NAME}.
     *
     * @param int $opUid      商户后台用户UID.
     * @param int ${NAME-LOW}Id {NAME-LOW}Id.
     *
     * @return mixed
     */
    public function doOffline($opUid, ${NAME-LOW}Id)
    {
        return \Service\Operation\{NAME}::instance()->doOffline($opUid, ${NAME-LOW}Id);
    }
}