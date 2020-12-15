<?php


namespace Module\{PARENT};

use Redis\RedisMultiCache;
use Util\TimeUtil;
use Util\Validator;

class {NAME} extends \Module\Base
{
    /**
     * 新增{NAME}.
     *
     * @param array $params 配置参数.
     *
     * @return bool|int|string
     * @throws \RpcBusinessException
     */
    public function create{NAME}(array $params)
    {
        //参数检查
        $params = $this->check{NAME}Params($params);

        $now = TimeUtil::now();
        //构建{NAME}数据
        ${NAME} = {BUILDSTR};
        //写入
        $db = \Model\Main\{NAME}::instance()->dbWrite();
        $db->beginTransaction();
        try {
            ${NAME-LOW}Id = \Model\Main\{NAME}::instance()->addNew(${NAME});
            if (empty(${NAME-LOW}Id)) {
                throw new \RpcBusinessException('{NAME}信息保存失败');
            }
            $db->commit();
            return ${NAME-LOW}Id;
        } catch (\Exception $e) {
            $db->rollback();
            throw $e;
        }
    }

    /**
     * 检查{NAME}新增/编辑参数.
     *
     * @param array $params 配置参数.
     * @param int ${NAME-LOW}Id {NAME-LOW}Id.
     * @return array
     * @throws \RpcBusinessException
     */
    private function check{NAME}Params(array $params, int ${NAME-LOW}Id = 0)
    {
        //检查{NAME}基本信息
        Validator::checkPositiveInteger(${NAME-LOW}Id, "参数错误");
        {CHECKSTR}
        return $params;
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
        ${NAME-LOW} = \Model\Main\{NAME}::instance()->getOneByCond([
                    'id' => ${NAME-LOW}Id,
                ], '*');
        if (empty(${NAME-LOW})) {
            throw new \RpcBusinessException('活动不存在');
        }
        return ${NAME-LOW};
    }

    /**
     * {NAME}列表.
     *
     * @param array $searchParams 搜索条件.
     * @param int $pageNo 分页页码.
     * @param int $pageSize 分页大小.
     *
     * @return array|\Db\JMDbConnectionPageResult
     * @throws \RpcBusinessException
     */
    public function get{NAME}List(array $searchParams, int $pageNo, int $pageSize)
    {
        $cond = [
            'is_delete' => 0,
        ];

        if (!empty($searchParams['name'])) {
            $condition['name LIKE'] = "%{$searchParams['name']}%";
        }
        $result = \Model\Main\{NAME}::instance()->getPageByCond($cond,"*",$'id DESC' ,$pageNo,$pageSize);
        if (empty($res->rows)) {
            return $res;
        }
        $now = TimeUtil::now();
        foreach ($result->rows as &$row) {
            $row['create_time'] = TimeUtil::fmtToTime($row['create_time']);
        }
        return $result;
    }

    /**
     * 删除{NAME}.
     *
     * @param int ${NAME-LOW}Id {NAME-LOW}Id.
     *
     * @return bool
     * @throws \RpcBusinessException
     */
    public function delete{NAME}(int ${NAME-LOW}Id)
    {
        ${NAME-LOW} = \Model\Main\{NAME}::instance()->getOneByCond([
            'id' => ${NAME-LOW}Id,
        ], '`id`, `is_delete`, `status`');

        if (empty(${NAME-LOW})) {
            throw new \RpcBusinessException('{NAME}不存在');
        }

        if (${NAME}['is_delete'] == 1) {
            //已经被删除的{NAME}不作处理
            return true;
        }     
        $now = TimeUtil::now();

        $db = \Model\Main\{NAME}::instance()->dbWrite();
        $db->beginTransaction();
        try {
            $res = \Model\Main\{NAME}::instance()->updateData([
                'is_delete' => \Enum\{NAME}::YES,
                'update_time' => $now,
            ], [
                'id' => ${NAME-LOW}Id,
            ]);
            if (empty($res)) {
                throw new \RpcBusinessException('删除{NAME}失败');
            }

            $db->commit();
            return true;
        } catch (\Exception $e) {
            $db->rollback();
            throw $e;
        }
    }

    /**
     * 更新{NAME}.
     *
     * @param int ${NAME-LOW}Id {NAME-LOW}Id.
     * @param array $params {NAME}配置参数.
     *
     * @return bool
     * @throws \RpcBusinessException
     */
    public function edit{NAME}(int ${NAME-LOW}Id, array $params)
    {
        ${NAME-LOW} = \Model\Main\{NAME}::instance()->getOneByCond([
            'id' => ${NAME-LOW}Id,
            'is_delete' => 0,
        ], '`id`');

        if (empty(${NAME-LOW})) {
            throw new \RpcBusinessException('{NAME}不存在');
        }

        //参数检查
        $params = $this->check{NAME}Params($params, ${NAME-LOW}Id);

        $now = TimeUtil::now();
        //构建{NAME}数据
        ${NAME-LOW} = {BUILDSTR};

        

        //写入
        $db = \Model\Main\{NAME}::instance()->dbWrite();
        $db->beginTransaction();
        try {
            //1.更新{NAME}基本数据
            $res = \Model\Main\{NAME}::instance()->updateData(${NAME-LOW}, [
                'id' => ${NAME-LOW}Id,
            ]);
            if (empty($res)) {
                throw new \RpcBusinessException('{NAME}信息更新失败');
            }

           
            $db->commit();
            return true;
        } catch (\Exception $e) {
            $db->rollback();
            throw $e;
        }
    }

    /**
     * 上线{NAME}.
     *
     * @param int ${NAME-LOW}Id {NAME-LOW}Id.
     *
     * @return bool
     * @throws \RpcBusinessException
     */
    public function doOnline(int ${NAME-LOW}Id)
    {
        ${NAME-LOW} = \Model\Main\{NAME}::instance()->getOneByCond([
            'id' => ${NAME-LOW}Id,
            'is_delete' => 0,
        ], 'id, status');

        if (empty(${NAME-LOW})) {
            throw new \RpcBusinessException('{NAME}不存在');
        }
        if (${NAME-LOW}['status'] == 1) {
            return true;
        }
        $update = [
            'status' => 1,
            'update_time' => $now,
        ];
        $res = \Model\Main\{NAME}::instance()->updateData($update, [
            'id' => ${NAME-LOW}Id,
        ]);
        return (bool)$res;
    }

    /**
     * 下线/结束{NAME}.
     *
     * @param int ${NAME-LOW}Id {NAME-LOW}Id.
     *
     * @return bool
     * @throws \RpcBusinessException
     */
    public function doOffline(int ${NAME-LOW}Id)
    {
        ${NAME-LOW} = \Model\Main\{NAME}::instance()->getOneByCond([
            'id' => ${NAME-LOW}Id,
            'is_delete' => \Enum\{NAME}::NO,
        ], '`id`, `status`,');

        if (empty(${NAME-LOW})) {
            throw new \RpcBusinessException('{NAME}不存在');
        }

        if (${NAME-LOW}['status'] == 0) {
            return true;
        }

        $now = TimeUtil::now();

        $update = [
            'status' => 0,
            'update_time' => $now,
        ];
        $res = \Model\Main\{NAME}::instance()->updateData($update, [
            'id' => ${NAME-LOW}Id,
        ]);
        return (bool)$res;
    }
}