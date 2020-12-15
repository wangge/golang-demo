<?php


namespace Module\{{.Parent}};

use Redis\RedisMultiCache;
use Util\TimeUtil;
use Util\Validator;

class {{.Name}} extends \Module\Base
{

    {{if .Per.Create}}
    /**
     * 新增{{.Table}}.
     *
     * @param array $params {{.Table}}参数.
     *
     * @return bool|int|string
     * @throws \RpcBusinessException
     */
    public function create{{.Name}}(array $params)
    {
        //参数检查
        $params = $this->check{{.Name}}Params($params);

        $now = TimeUtil::now();
        //构建{{.Name}}数据
        ${{.NameLow}} = [
                    {{range .Fd}}
                    '{{.FieldName}}' => $params['{{.FieldName}}'],
                    {{end}}
                ];
        //写入
        $db = \Model\Main\{{.Name}}::instance()->dbWrite();
        $db->beginTransaction();
        try {
            ${{.NameLow}}Id = \Model\Main\{{.Name}}::instance()->addNew(${{.NameLow}});
            if (empty(${{.NameLow}}Id)) {
                throw new \RpcBusinessException('信息保存失败');
            }
            $db->commit();
            return ${{.NameLow}}Id;
        } catch (\Exception $e) {
            $db->rollback();
            throw $e;
        }
    }
    {{end}}
    {{if or .Per.Create .Per.Edit}}
    /**
     * 检查{{.Table}}新增/编辑参数.
     *
     * @param array $params {{.Table}}参数.
     * @param int ${{.NameLow}}Id {{.Table}}Id.
     * @return array
     * @throws \RpcBusinessException
     */
    private function check{{.Name}}Params(array $params, int ${{.NameLow}}Id = 0)
    {

        {{range .Fd}}
        Validator::checkEmpty($params['{{.FieldName}}'], "{{.FieldDesc}}不能为空");
        {{end}}
        // 检测{{.Table}}是否重复
        ${{.NameLow}} = \Model\Main\{{.Name}}::instance()->getOneByCond([
            'name' => $params['name'],
            'is_delete' => 0
        ],'id');
        if(!empty(${{.NameLow}}) && ${{.NameLow}}['id'] != ${{.NameLow}}Id)
        {
            throw new \RpcBusinessException('{{.Table}}名称已经存在');
        }
        return $params;
    }
    {{end}}
    {{if .Per.Detail}}
    /**
     * {{.Table}}详情.
     *
     * @param int ${{.NameLow}}Id {{.Table}}Id.
     *
     * @return array
     * @throws \RpcBusinessException
     */
    public function get{{.Name}}Detail(int ${{.NameLow}}Id)
    {
        ${{.NameLow}} = \Model\Main\{{.Name}}::instance()->getOneByCond([
                    'id' => ${{.NameLow}}Id,
                ], '*');
        if (empty(${{.NameLow}})) {
            throw new \RpcBusinessException('{{.Table}}不存在');
        }
        return ${{.NameLow}};
    }
{{end}}
{{if .Per.List}}
    /**
     * {{.Table}}列表.
     *
     * @param array $searchParams 搜索条件.
     * @param int $pageNo 分页页码.
     * @param int $pageSize 分页大小.
     *
     * @return array|\Db\JMDbConnectionPageResult
     * @throws \RpcBusinessException
     */
    public function get{{.Name}}List(array $searchParams, int $pageNo, int $pageSize)
    {
        $cond = [
            'is_delete' => 0,
        ];
        {{range .Fd}}
        '{{.FieldName}}' => $params['{{.FieldName}}'],
        if (isset($searchParams['{{.FieldName}}']) && !empty($searchParams['{{.FieldName}}'])) {
            $cond['{{.FieldName}}'] = "%{$searchParams['{{.FieldName}}']}%";
        }
        {{end}}
        if (isset($searchParams['name']) && !empty($searchParams['name'])) {
            $cond['name LIKE'] = "%{$searchParams['name']}%";
        }
        $result = \Model\Main\{{.Name}}::instance()->getPageByCond($cond,"*",'id DESC' ,$pageNo,$pageSize);
        if (empty($result->rows)) {
            return [];
        }
        $now = TimeUtil::now();
        foreach ($result->rows as &$row) {
            $row['create_time'] = TimeUtil::fmtToTime($row['create_time']);
        }
        return $result;
    }
{{end}}
{{if .Per.Delete}}
    /**
     * 删除{{.Table}}.
     *
     * @param int ${{.NameLow}}Id {{.Table}}Id.
     *
     * @return bool
     * @throws \RpcBusinessException
     */
    public function delete{{.Name}}(int ${{.NameLow}}Id)
    {
        ${{.NameLow}} = \Model\Main\{{.Name}}::instance()->getOneByCond([
            'id' => ${{.NameLow}}Id,
        ], '`id`, `is_delete`, `status`');

        if (empty(${{.NameLow}})) {
            throw new \RpcBusinessException('{{.Table}}不存在');
        }

        if (${{.NameLow}}['is_delete'] == 1) {
            //已经被删除的{{.Table}}不作处理
            return true;
        }     
        $now = TimeUtil::now();

        $db = \Model\Main\{{.Name}}::instance()->dbWrite();
        $db->beginTransaction();
        try {
            $res = \Model\Main\{{.Name}}::instance()->updateData([
                'is_delete' => 1,
                'update_time' => $now,
            ], [
                'id' => ${{.NameLow}}Id,
            ]);
            if (empty($res)) {
                throw new \RpcBusinessException('删除{{.Table}}失败');
            }

            $db->commit();
            return true;
        } catch (\Exception $e) {
            $db->rollback();
            throw $e;
        }
    }
{{end}}
{{if .Per.Edit}}
    /**
     * 更新{{.Table}}.
     *
     * @param int ${{.NameLow}}Id {{.Table}}Id.
     * @param array $params {{.Table}}参数.
     *
     * @return bool
     * @throws \RpcBusinessException
     */
    public function edit{{.Name}}(int ${{.NameLow}}Id, array $params)
    {
        ${{.NameLow}} = \Model\Main\{{.Name}}::instance()->getOneByCond([
            'id' => ${{.NameLow}}Id,
            'is_delete' => 0,
        ], '`id`');

        if (empty(${{.NameLow}})) {
            throw new \RpcBusinessException('{{.Table}}不存在');
        }

        //参数检查
        $params = $this->check{{.Name}}Params($params, ${{.NameLow}}Id);

        $now = TimeUtil::now();
        //构建{{.Name}}数据
        ${{.NameLow}} = [
            {{range .Fd}}
            '{{.FieldName}}' => $params['{{.FieldName}}'],
            {{end}}
        ];

        

        //写入
        $db = \Model\Main\{{.Name}}::instance()->dbWrite();
        $db->beginTransaction();
        try {
            //1.更新{{.Table}}基本数据
            $res = \Model\Main\{{.Name}}::instance()->updateData(${{.NameLow}}, [
                'id' => ${{.NameLow}}Id,
            ]);
            if (empty($res)) {
                throw new \RpcBusinessException('{{.Table}}信息更新失败');
            }

           
            $db->commit();
            return true;
        } catch (\Exception $e) {
            $db->rollback();
            throw $e;
        }
    }
{{end}}
{{if .Per.Online}}
    /**
     * 上线{{.Table}}.
     *
     * @param int ${{.NameLow}}Id {{.Table}}Id.
     *
     * @return bool
     * @throws \RpcBusinessException
     */
    public function doOnline(int ${{.NameLow}}Id)
    {
        ${{.NameLow}} = \Model\Main\{{.Name}}::instance()->getOneByCond([
            'id' => ${{.NameLow}}Id,
            'is_delete' => 0,
        ], 'id, status');

        if (empty(${{.NameLow}})) {
            throw new \RpcBusinessException('{{.Table}}不存在');
        }
        if (${{.NameLow}}['status'] == 1) {
            return true;
        }
        $now = TimeUtil::now();
        $update = [
            'status' => 1,
            'update_time' => $now,
        ];
        $res = \Model\Main\{{.Name}}::instance()->updateData($update, [
            'id' => ${{.NameLow}}Id,
        ]);
        return (bool)$res;
    }
{{end}}
{{if .Per.Offline}}
    /**
     * 下线/结束{{.Table}}.
     *
     * @param int ${{.NameLow}}Id {{.NameLow}}Id.
     *
     * @return bool
     * @throws \RpcBusinessException
     */
    public function doOffline(int ${{.NameLow}}Id)
    {
        ${{.NameLow}} = \Model\Main\{{.Name}}::instance()->getOneByCond([
            'id' => ${{.NameLow}}Id,
            'is_delete' => 0,
        ], '`id`, `status`');

        if (empty(${{.NameLow}})) {
            throw new \RpcBusinessException('{{.Table}}不存在');
        }

        if (${{.NameLow}}['status'] == 0) {
            return true;
        }

        $now = TimeUtil::now();

        $update = [
            'status' => 0,
            'update_time' => $now,
        ];
        $res = \Model\Main\{{.Name}}::instance()->updateData($update, [
            'id' => ${{.NameLow}}Id,
        ]);
        return (bool)$res;
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

        $result = \Model\Main\{{.Name}}::instance()->getAllByCond();
        if(empty($result))
                {
                    throw new \RpcBusinessException("没有可导出的数据");
                }
                foreach($result as &$row)
                {
                    $row['sale_price'] = \Util\MoneyHelper::fmtFenToYuan($row['sale_price']);
                }
                $fields = [
                        {{range .Fd}}
                            ['key' => '{{.FieldName}}',      'name'=>'{{.FieldDesc}}'],
                        {{end}}

                ];
                $title = "{{.Name}}";
                $fileUrl = ExportUtil::exportExcel($fields, $result, $title);
                $fileUrl = OssUtil::signOssPrivateFileUrl($fileUrl, 1800, $title.'['.TimeUtil::fmtToTime(TimeUtil::now()).'].xls');

                return $fileUrl;
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
         * @throws \RpcBusinessException
         */
        public function audit(int ${{.NameLow}}Id, $params)
        {
            ${{.NameLow}} = \Model\Main\{{.Name}}::instance()->getOneByCond([
                'id' => ${{.NameLow}}Id,
                'is_delete' => 0,
            ], 'id, status');

            if (empty(${{.NameLow}})) {
                throw new \RpcBusinessException('{{.Table}}不存在');
            }
            // todo 处理审核
        }
    {{end}}

    {{range .Fcs}}
    /**
        *{{.Des}}.
     */
     public function {{.Fname}}($opUid{{range .Param}} ,${{.}} {{end}})
     {

     }
    {{end}}
}