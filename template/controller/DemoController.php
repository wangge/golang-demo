<?php

/**
 * {NAME}管理
 */
class Controller_Base_{NAME} extends Controller_Base_Base
{

    /**
     * 创建{NAME}
     */
    public function action_Create{NAME}()
    {
        $params = {PARAMS};

        $result = \Module\{NAME}::instance()->create{NAME}($this->uid, $params);
        if (empty($result)) {
            $this->responseFail('保存失败');
        }
        $this->responseSuccess(['{NAME-LOW}Id' => intval($result)], '保存成功');
    }

    /**
     * 活动详情
     */
    public function action_{NAME}Detail()
    {
        ${NAME-LOW}Id = ZKGetPostInt('{NAME-LOW}_id', 0);
        if (${NAME-LOW}Id <= 0) {
            $this->responseFail('参数错误');
        }

        ${NAME-LOW} = \Module\{NAME}::instance()->get{NAME}Detail(${NAME-LOW}Id);
        if (empty(${NAME-LOW})) {
            $this->responseFail('{NAME}数据不存在');
        }
        $this->responseSuccess(${NAME-LOW}, '数据获取成功');
    }

    /**
     * {NAME}列表
     */
    public function action_{NAME}List()
    {
        $searchParams = [
            'name'   => ZKGetPost('name', ''),
            'status' => ZKGetPost('status', ''),
        ];
        $pageNo = ZKGetPostInt('page_no', 1);
        $pageSize = ZKGetPostInt('page_size', 10);

        $result = \Module\{NAME}::instance()->get{NAME}List($searchParams, $pageNo, $pageSize);
        $this->responseSuccess($result, '数据获取成功');
    }

    /**
     * 删除{NAME}
     */
    public function action_Delete()
    {
        ${NAME-LOW}Id = ZKGetPostInt('{NAME-LOW}_id', 0);
        if (${NAME-LOW}Id <= 0) {
            $this->responseFail('参数错误');
        }

        $res = \Module\{NAME}::instance()->delete{NAME}($this->uid, ${NAME-LOW}Id);
        if (empty($res)) {
            $this->responseFail('操作失败');
        }
        $this->responseSuccess(true, '操作成功');
    }

    /**
     * 编辑{NAME}
     */
    public function action_Edit{NAME}()
    {
        ${NAME-LOW}Id = ZKGetPostInt('{NAME-LOW}_id', 0);
        if (${NAME-LOW}Id <= 0) {
            $this->responseFail('参数错误');
        }

        $params = {PARAMS};

        $result = \Module\{NAME}::instance()->edit{NAME}($this->uid, ${NAME-LOW}Id, $params);
        if (empty($result)) {
            $this->responseFail('编辑失败');
        }
        $this->responseSuccess(boolval($result), '编辑成功');
    }

    /**
     * 上线{NAME}
     */
    public function action_DoOnline()
    {
        ${NAME-LOW}Id = ZKGetPostInt('{NAME-LOW}_id', 0);
        if (${NAME-LOW}Id <= 0) {
            $this->responseFail('参数错误');
        }

        $res = \Module\{NAME}::instance()->doOnline($this->uid, ${NAME-LOW}Id);
        if ($res) {
            $this->responseSuccess($res, '更新成功');
        } else {
            $this->responseFail('更新失败');
        }
    }

    /**
     * 下线{NAME}
     */
    public function action_DoOffline()
    {
        ${NAME-LOW}Id = ZKGetPostInt('{NAME-LOW}_id', 0);
        if (${NAME-LOW}Id <= 0) {
            $this->responseFail('参数错误');
        }

        $res = \Module\{NAME}::instance()->doOffline($this->uid, ${NAME-LOW}Id);
        if ($res) {
            $this->responseSuccess($res, '更新成功');
        } else {
            $this->responseFail('更新失败');
        }
    }
}