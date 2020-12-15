<?php
namespace Handler\{PARENT};

class {NAME} extends \Handler\Base
{
    /**
     * 新增{NAME}.
     *
     * @param int   $opUid  操作人员UID.
     * @param array $params 配置参数.
     *
     * @return bool|int|string
     * @throws \Exception
     */
    public function create{NAME}(int $opUid, array $params)
    {
        $serviceLogger = \Util\ServiceLogger::getNew(__CLASS__, __FUNCTION__, func_get_args());
        $serviceLogger->setUid($opUid);
        $serviceLogger->setKeyword(__FUNCTION__);
        try {
            $serviceLogger->logRequest();
            $response = \Module\{PARENT}\{NAME}::instance()->create{NAME}($params);
            $serviceLogger->logSuccess($response);
            return $response;
        } catch (\Exception $e) {
            $serviceLogger->logException($e);
            throw $e;
        }
    }

    /**
     * {NAME}列表.
     *
     * @param array $searchParams 搜索条件.
     * @param int   $pageNo       分页页码.
     * @param int   $pageSize     分页大小.
     *
     * @return array|\Db\JMDbConnectionPageResult
     * @throws \RpcBusinessException
     */
    public function get{NAME}List(array $searchParams = [], int $pageNo = 1, int $pageSize = 10)
    {
        return \Module\{PARENT}\{NAME}::instance()->get{NAME}List($searchParams, $pageNo, $pageSize);
    }

    /**
     * {NAME}详情.
     *
     * @param int ${NAME-LOW}Id {NAME-LOW}Id.
     *
     * @return array
     * @throws \RpcBusinessException
     */
    public function get{NAME}Detail(int ${NAME-LOW}Id)
    {
        return \Module\{PARENT}\{NAME}::instance()->get{NAME}Detail(${NAME-LOW}Id);
    }

    /**
     * 删除{NAME}.
     *
     * @param int $opUid      操作人员UID.
     * @param int ${NAME-LOW}Id {NAME-LOW}Id.
     *
     * @return bool
     * @throws \Exception
     */
    public function delete{NAME}(int $opUid, int ${NAME-LOW}Id)
    {
        $serviceLogger = \Util\ServiceLogger::getNew(__CLASS__, __FUNCTION__, func_get_args());
        $serviceLogger->setUid($opUid);
        $serviceLogger->setKeyword(__FUNCTION__);
        try {
            $serviceLogger->logRequest();
            $response = \Module\{PARENT}\{NAME}::instance()->delete{NAME}(${NAME-LOW}Id);
            $serviceLogger->logSuccess($response);
            return $response;
        } catch (\Exception $e) {
            $serviceLogger->logException($e);
            throw $e;
        }
    }

    /**
     * 编辑{NAME}.
     *
     * @param int   $opUid      操作用户UID.
     * @param int   ${NAME-LOW}Id {NAME-LOW}Id.
     * @param array $params     {NAME}配置参数.
     *
     * @return mixed
     * @throws \Exception
     */
    public function edit{NAME}(int $opUid, int ${NAME-LOW}Id, array $params)
    {
        $serviceLogger = \Util\ServiceLogger::getNew(__CLASS__, __FUNCTION__, func_get_args());
        $serviceLogger->setUid($opUid);
        $serviceLogger->setKeyword(__FUNCTION__);
        try {
            $serviceLogger->logRequest();
            $response = \Module\{PARENT}\{NAME}::instance()->edit{NAME}(${NAME-LOW}Id, $params);
            $serviceLogger->logSuccess($response);
            return $response;
        } catch (\Exception $e) {
            $serviceLogger->logException($e);
            throw $e;
        }
    }

    /**
     * 上线{NAME}.
     *
     * @param int $opUid      操作人员UID.
     * @param int ${NAME-LOW}Id {NAME-LOW}Id.
     *
     * @return bool
     * @throws \Exception
     */
    public function doOnline(int $opUid, int ${NAME-LOW}Id)
    {
        $serviceLogger = \Util\ServiceLogger::getNew(__CLASS__, __FUNCTION__, func_get_args());
        $serviceLogger->setUid($opUid);
        $serviceLogger->setKeyword(__FUNCTION__);
        try {
            $serviceLogger->logRequest();
            $response = \Module\{PARENT}\{NAME}::instance()->doOnline(${NAME-LOW}Id);
            $serviceLogger->logSuccess($response);
            return $response;
        } catch (\Exception $e) {
            $serviceLogger->logException($e);
            throw $e;
        }
    }

    /**
     * 下线/结束{NAME}.
     *
     * @param int $opUid      操作人员UID.
     * @param int ${NAME-LOW}Id {NAME-LOW}Id.
     *
     * @return bool
     * @throws \Exception
     */
    public function doOffline(int $opUid, int ${NAME-LOW}Id)
    {
        $serviceLogger = \Util\ServiceLogger::getNew(__CLASS__, __FUNCTION__, func_get_args());
        $serviceLogger->setUid($opUid);
        $serviceLogger->setKeyword(__FUNCTION__);
        try {
            $serviceLogger->logRequest();
            $response = \Module\{PARENT}\{NAME}::instance()->doOffline(${NAME-LOW}Id);
            $serviceLogger->logSuccess($response);
            return $response;
        } catch (\Exception $e) {
            $serviceLogger->logException($e);
            throw $e;
        }
    }
}