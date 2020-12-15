<?php
/**
 * {{.Table}}管理
 */
class Controller_Base_{{.Name}} extends Controller_Base_Base
{
    {{if .Per.Create}}
    /**
     * 创建{{.Table}}
     */
    public function action_Create()
    {
        $params = [{{range .Fd}}
        '{{.FieldName}}' => ZKGetPost('{{.FieldName}}',''),//{{.FieldDesc}} {{end}}
        ];
        $result = \Module\{{.Name}}::instance()->create{{.Name}}($this->uid, $params);
        if (empty($result)) {
            $this->responseFail('保存失败');
        }
        $this->responseSuccess(['{{.NameLow}}Id' => intval($result)], '保存成功');
    }
    {{end}}
    {{if .Per.Detail}}
    /**
     * 查询{{.Table}}详情
     */
    public function action_Detail()
    {
        ${{.NameLow}}Id = ZKGetPostInt('{{.NameLow}}_id', 0);
        if (${{.NameLow}}Id <= 0) {
            $this->responseFail('参数错误');
        }
        ${{.NameLow}} = \Module\{{.Name}}::instance()->get{{.Name}}Detail(${{.NameLow}}Id);
        if (empty(${{.NameLow}})) {
            $this->responseFail('{{.Table}}数据不存在');
        }
        $this->responseSuccess(${{.NameLow}}, '数据获取成功');
    }
    {{end}}
    {{if .Per.List}}
    /**
     * 查询{{.Table}}列表
     */
    public function action_List()
    {
        $searchParams = [{{range .Search}}
        '{{.SName}}' => ZKGetPost('{{.SName}}',''),// {{.SDes}}{{end}}
         ];
        $pageNo = ZKGetPostInt('page_no', 1);
        $pageSize = ZKGetPostInt('page_size', 10);
        $result = \Module\{{.Name}}::instance()->get{{.Name}}List($searchParams, $pageNo, $pageSize);
        $this->responseSuccess($result, '数据获取成功');
    }
    {{end}}
    {{if .Per.Delete}}
    /**
     * 删除{{.Table}}
     */
    public function action_Delete()
    {
        ${{.NameLow}}Id = ZKGetPostInt('{{.NameLow}}_id', 0);
        if (${{.NameLow}}Id <= 0) {
            $this->responseFail('参数错误');
        }

        $res = \Module\{{.Name}}::instance()->delete{{.Name}}($this->uid, ${{.NameLow}}Id);
        if (empty($res)) {
            $this->responseFail('删除失败');
        }
        $this->responseSuccess(true, '删除成功');
    }
    {{end}}
    {{if .Per.Edit}}
    /**
     * 编辑{{.Table}}
     */
    public function action_Edit()
    {
        ${{.NameLow}}Id = ZKGetPostInt('{{.NameLow}}_id', 0);
        if (${{.NameLow}}Id <= 0) {
            $this->responseFail('参数错误');
        }
        $params = [{{range .Fd}}
                '{{.FieldName}}' => ZKGetPost('{{.FieldName}}',''),//{{.FieldDesc}}{{end}}
                ];
        $result = \Module\{{.Name}}::instance()->edit{{.Name}}($this->uid, ${{.NameLow}}Id, $params);
        if (empty($result)) {
            $this->responseFail('编辑失败');
        }
        $this->responseSuccess(boolval($result), '编辑成功');
    }
    {{end}}
    {{if .Per.Online}}
    /**
     * 上线{{.Table}}
     */
    public function action_DoOnline()
    {
        ${{.NameLow}}Id = ZKGetPostInt('{{.NameLow}}_id', 0);
        if (${{.NameLow}}Id <= 0) {
            $this->responseFail('参数错误');
        }

        $res = \Module\{{.Name}}::instance()->doOnline($this->uid, ${{.NameLow}}Id);
        if ($res) {
            $this->responseSuccess($res, '更新成功');
        } else {
            $this->responseFail('更新失败');
        }
    }
    {{end}}

    {{if .Per.Offline}}
    /**
     * 下线{{.Table}}
     */
    public function action_DoOffline()
    {
        ${{.NameLow}}Id = ZKGetPostInt('{{.NameLow}}_id', 0);
        if (${{.NameLow}}Id <= 0) {
            $this->responseFail('参数错误');
        }

        $res = \Module\{{.Name}}::instance()->doOffline($this->uid, ${{.NameLow}}Id);
        if ($res) {
            $this->responseSuccess($res, '更新成功');
        } else {
            $this->responseFail('更新失败');
        }
    }
    {{end}}

    {{if .Per.Export}}
    /**
     * 导出{{.Table}}列表.
     */
    public function action_ExportList()
    {
        $searchParams = [{{range .Search}}
        '{{.SName}}' => ZKGetPost('{{.SName}}',''),// {{.SDes}}{{end}}
         ];
        $fileUrl = \Module\{{.Name}}::instance()->export{{.Name}}List($searchParams);
        if (!empty($fileUrl)) {
            $this->responseSuccess([
                'file_url' => $fileUrl,
            ], '数据导出成功');
        } else {
            $this->responseFail('数据导出失败');
        }
    }
    {{end}}
    {{if .Per.Audit}}
    /**
     * 审核{{.Table}}.
     */
    public function action_Audit()
    {
        ${{.NameLow}}Id = ZKGetPostInt('{{.NameLow}}_id', 0);
        if (${{.NameLow}}Id <= 0) {
            $this->responseFail('参数错误');
        }
        $params = [
        "status" => ZKGetPostInt('status',''),
        "remark" => ZKGetPost('remark',''),
        ];
        $res = \Module\{{.Name}}::instance()->audit($this->uid, ${{.NameLow}}Id, $params);
        if ($res) {
            $this->responseSuccess($res, '审核成功');
        } else {
            $this->responseFail('审核失败');
        }
    }
    {{end}}
    {{range .Fcs}}
    /**
    *{{.Des}}.
    */
    public function action_{{.Fname}}()
    {
        {{if .NeedPk}}
        ${{.PK}} = ZKGetPost("{{.PK}}");{{end}}
        {{if .ParamIsArr}}
        $params = [{{range .Param}}
        '{{.}}' => ZKGetPost("{{.}}"),{{end}}
        ];
        $res = \Module\{{$.Name}}::instance()->{{.Fname}}($this->uid,{{if .NeedPk}} ${{.PK}},{{end}} $params);
        {{else}}{{range .Param}}
        ${{.}}= ZKGetPost("{{.}}");
        {{end}}
        $res = \Module\{{$.Name}}::instance()->{{.Fname}}($this->uid {{range .Param}} ,${{.}} {{end}});
        {{end}}
        if ($res) {
        $this->responseSuccess($res, '操作成功');
        } else {
        $this->responseFail('操作失败');
        }
    }
    {{end}}
}