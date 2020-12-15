<?php


namespace Module;


class {{.Name}} extends ModuleBase
{
    {{if .Per.Create}}
    /**
     * 创建{{.Table}}.
     *
     * @param int   $opUid  操作用户UID.
     * @param array $params {{.Table}}参数.
     *
     * @return mixed
     */
    public function create{{.Name}}(int $opUid, array $params)
    {
        return \Service\ZKService\{{.Name}}::instance()->create{{.Name}}($opUid, $params);
    }
    {{end}}
    {{if .Per.Detail}}
    /**
     * {{.Table}}详情.
     *
     * @param int ${{.NameLow}}Id {{.Table}}ID.
     *
     * @return mixed
     */
    public function get{{.Name}}Detail(int ${{.NameLow}}Id)
    {
        return \Service\ZKService\{{.Name}}::instance()->get{{.Name}}Detail(${{.NameLow}}Id);
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
     * @return mixed
     */
    public function get{{.Name}}List(array $searchParams = [], $pageNo = 1, $pageSize = 10)
    {
        return \Service\ZKService\{{.Name}}::instance()->get{{.Name}}List($searchParams, $pageNo, $pageSize);
    }
{{end}}
   {{if .Per.Delete}}
    /**
     * 删除{{.Table}}.
     *
     * @param int $opUid      操作用户UID.
     * @param int ${{.NameLow}}Id {{.Table}}ID.
     *
     * @return mixed
     */
    public function delete{{.Name}}(int $opUid, int ${{.NameLow}}Id)
    {
        return \Service\ZKService\{{.Name}}::instance()->delete{{.Name}}($opUid, ${{.NameLow}}Id);
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
     */
    public function edit{{.Name}}(int $opUid, int ${{.NameLow}}Id, array $params)
    {
        return \Service\ZKService\{{.Name}}::instance()->edit{{.Name}}($opUid, ${{.NameLow}}Id, $params);
    }
{{end}}
    {{if .Per.Online}}
    /**
     * 上线{{.Table}}.
     *
     * @param int $opUid      商户后台用户UID.
     * @param int ${{.NameLow}}Id {{.Table}}Id.
     *
     * @return mixed
     */
    public function doOnline($opUid, ${{.NameLow}}Id)
    {
        return \Service\ZKService\{{.Name}}::instance()->doOnline($opUid, ${{.NameLow}}Id);
    }
{{end}}
    {{if .Per.Offline}}
    /**
     * 下线/结束{{.Table}}.
     *
     * @param int $opUid      商户后台用户UID.
     * @param int ${{.NameLow}}Id {{.Table}}Id.
     *
     * @return mixed
     */
    public function doOffline($opUid, ${{.NameLow}}Id)
    {
        return \Service\ZKService\{{.Name}}::instance()->doOffline($opUid, ${{.NameLow}}Id);
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
        return \Service\ZKService\{{.Name}}::instance()->export{{.Name}}List($searchParams);
    }
    {{end}}

    {{if .Per.Audit}}
    /**
     * 审核{{.Table}}
     *
     * @param int $opUid      商户后台用户UID.
     * @param int ${{.NameLow}}Id {{.Table}}Id.
     * @param array $params .
     * @return mixed
     */
    public function audit($opUid, ${{.NameLow}}Id, $params)
    {
        return \Service\ZKService\{{.Name}}::instance()->audit($opUid, ${{.NameLow}}Id, $params);
    }
    {{end}}

    {{range .Fcs}}
    /**
        *{{.Des}}.
     */
     public function {{.Fname}}($opUid{{if .NeedPk}} , ${{.PK}}{{end}}{{if .ParamIsArr}}, $params {{else}} {{range .Param}} ,${{.}} {{end}} {{end}})
     {
        return \Service\ZKService\{{$.Name}}::instance()->{{.Fname}}($opUid{{if .NeedPk}} , ${{.PK}}{{end}}{{if .ParamIsArr}}, $params {{else}} {{range .Param}} ,${{.}} {{end}} {{end}});
     }
    {{end}}
}
