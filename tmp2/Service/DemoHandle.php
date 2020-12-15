<?php
namespace Handler\{{.Parent}};

class {{.Name}} extends \Handler\Base
{
   {{if .Per.Create}}
    /**
     * 新增{{.Table}}.
     *
     * @param int   $opUid  操作人员UID.
     * @param array $params {{.Table}}参数.
     *
     * @return bool|int|string
     * @throws \Exception
     */
    public function create{{.Name}}(int $opUid, array $params)
    {
        $serviceLogger = \Util\ServiceLogger::getNew(__CLASS__, __FUNCTION__, func_get_args());
        $serviceLogger->setUid($opUid);
        $serviceLogger->setKeyword(__FUNCTION__);
        try {
            $serviceLogger->logRequest();
            $response = \Module\{{.Parent}}\{{.Name}}::instance()->create{{.Name}}($params);
            $serviceLogger->logSuccess($response);
            return $response;
        } catch (\Exception $e) {
            $serviceLogger->logException($e);
            throw $e;
        }
    }
{{end}}
    {{if .Per.List}}
    /**
     * {{.Table}}列表.
     *
     * @param array $searchParams 搜索条件.
     * @param int   $pageNo       分页页码.
     * @param int   $pageSize     分页大小.
     *
     * @return array|\Db\JMDbConnectionPageResult
     * @throws \RpcBusinessException
     */
    public function get{{.Name}}List(array $searchParams = [], int $pageNo = 1, int $pageSize = 10)
    {
        return \Module\{{.Parent}}\{{.Name}}::instance()->get{{.Name}}List($searchParams, $pageNo, $pageSize);
    }
{{end}}
    {{if .Per.Detail}}
    /**
     * {{.Table}}详情.
     *
     * @param int ${{.NameLow}}Id {{.NameLow}}Id.
     *
     * @return array
     * @throws \RpcBusinessException
     */
    public function get{{.Name}}Detail(int ${{.NameLow}}Id)
    {
        return \Module\{{.Parent}}\{{.Name}}::instance()->get{{.Name}}Detail(${{.NameLow}}Id);
    }
{{end}}
     {{if .Per.Delete}}
    /**
     * 删除{{.Table}}.
     *
     * @param int $opUid      操作人员UID.
     * @param int ${{.NameLow}}Id {{.Table}}Id.
     *
     * @return bool
     * @throws \Exception
     */
    public function delete{{.Name}}(int $opUid, int ${{.NameLow}}Id)
    {
        $serviceLogger = \Util\ServiceLogger::getNew(__CLASS__, __FUNCTION__, func_get_args());
        $serviceLogger->setUid($opUid);
        $serviceLogger->setKeyword(__FUNCTION__);
        try {
            $serviceLogger->logRequest();
            $response = \Module\{{.Parent}}\{{.Name}}::instance()->delete{{.Name}}(${{.NameLow}}Id);
            $serviceLogger->logSuccess($response);
            return $response;
        } catch (\Exception $e) {
            $serviceLogger->logException($e);
            throw $e;
        }
    }
{{end}}
    {{if .Per.Edit}}
    /**
     * 编辑{{.Table}}.
     *
     * @param int   $opUid      操作用户UID.
     * @param int   ${{.NameLow}}Id {{.Table}}Id.
     * @param array $params     {{.Table}}参数.
     *
     * @return mixed
     * @throws \Exception
     */
    public function edit{{.Name}}(int $opUid, int ${{.NameLow}}Id, array $params)
    {
        $serviceLogger = \Util\ServiceLogger::getNew(__CLASS__, __FUNCTION__, func_get_args());
        $serviceLogger->setUid($opUid);
        $serviceLogger->setKeyword(__FUNCTION__);
        try {
            $serviceLogger->logRequest();
            $response = \Module\{{.Parent}}\{{.Name}}::instance()->edit{{.Name}}(${{.NameLow}}Id, $params);
            $serviceLogger->logSuccess($response);
            return $response;
        } catch (\Exception $e) {
            $serviceLogger->logException($e);
            throw $e;
        }
    }
{{end}}
    {{if .Per.Online}}
    /**
     * 上线{{.Table}}.
     *
     * @param int $opUid      操作人员UID.
     * @param int ${{.NameLow}}Id {{.Table}}Id.
     *
     * @return bool
     * @throws \Exception
     */
    public function doOnline(int $opUid, int ${{.NameLow}}Id)
    {
        $serviceLogger = \Util\ServiceLogger::getNew(__CLASS__, __FUNCTION__, func_get_args());
        $serviceLogger->setUid($opUid);
        $serviceLogger->setKeyword(__FUNCTION__);
        try {
            $serviceLogger->logRequest();
            $response = \Module\{{.Parent}}\{{.Name}}::instance()->doOnline(${{.NameLow}}Id);
            $serviceLogger->logSuccess($response);
            return $response;
        } catch (\Exception $e) {
            $serviceLogger->logException($e);
            throw $e;
        }
    }
{{end}}
    {{if .Per.Offline}}
    /**
     * 下线/结束{{.Table}}.
     *
     * @param int $opUid      操作人员UID.
     * @param int ${{.NameLow}}Id {{.Table}}Id.
     *
     * @return bool
     * @throws \Exception
     */
    public function doOffline(int $opUid, int ${{.NameLow}}Id)
    {
        $serviceLogger = \Util\ServiceLogger::getNew(__CLASS__, __FUNCTION__, func_get_args());
        $serviceLogger->setUid($opUid);
        $serviceLogger->setKeyword(__FUNCTION__);
        try {
            $serviceLogger->logRequest();
            $response = \Module\{{.Parent}}\{{.Name}}::instance()->doOffline(${{.NameLow}}Id);
            $serviceLogger->logSuccess($response);
            return $response;
        } catch (\Exception $e) {
            $serviceLogger->logException($e);
            throw $e;
        }
    }
{{end}}
    {{if .Per.Export}}
    /**
     * 导出{{.Table}}列表
     *
     * @param array $searchParams 搜索条件.
     * @return mixed
     */
    public function export{{.Name}}List($searchParams)
    {
        return \Module\{{.Parent}}\{{.Name}}::instance()->export{{.Name}}List($searchParams);
    }
    {{end}}

    {{if .Per.Audit}}
    /**
     * 审核{{.Table}}
      *
      * @param int $opUid      商户后台用户UID.
      * @param int ${{.NameLow}}Id {{.Table}}Id.
      * @param array $params .
     * @return bool
     * @throws \Exception
     */
    public function audit(int $opUid, int ${{.NameLow}}Id, $params)
    {
        $serviceLogger = \Util\ServiceLogger::getNew(__CLASS__, __FUNCTION__, func_get_args());
        $serviceLogger->setUid($opUid);
        $serviceLogger->setKeyword(__FUNCTION__);
        try {
            $serviceLogger->logRequest();
            $response = \Module\{{.Parent}}\{{.Name}}::instance()->audit(${{.NameLow}}Id,$params);
            $serviceLogger->logSuccess($response);
            return $response;
        } catch (\Exception $e) {
            $serviceLogger->logException($e);
            throw $e;
        }
    }
    {{end}}

    {{range .Fcs}}
    /**
        *{{.Des}}.
    */
    public function {{.Fname}}($opUid{{if .NeedPk}} ,${{.PK}}{{end}} {{if .ParamIsArr}}, $params {{else}} {{range .Param}} ,${{.}} {{end}} {{end}})
    {
        $serviceLogger = \Util\ServiceLogger::getNew(__CLASS__, __FUNCTION__, func_get_args());
            $serviceLogger->setUid($opUid);
            $serviceLogger->setKeyword(__FUNCTION__);
            try {
                $serviceLogger->logRequest();
                $response = \Module\{{$.Parent}}\{{$.Name}}::instance()->{{.Fname}}({{if .NeedPk}} ${{.PK}}{{end}}{{if .ParamIsArr}}, $params {{else}} {{range .Param}} ${{.}} {{end}} {{end}});
                $serviceLogger->logSuccess($response);
                return $response;
            } catch (\Exception $e) {
                $serviceLogger->logException($e);
                throw $e;
            }
    }
    {{end}}
}