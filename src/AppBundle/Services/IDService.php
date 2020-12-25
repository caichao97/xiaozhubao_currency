<?php

namespace AppBundle\Services;

class IDService
{
    const EPOCH_TIME = 1540828800000; //时间基数

    const BITS_FULL     = 64;
    const BITS_PRE      = 2;  // 固定位01
    const BITS_TIME     = 41; // 可支持69年
    const BITS_SERVER   = 4;  // 可支持16台集群服务
    const BITS_WORKER   = 10;  // 可支持业务数1024个
    const BITS_SEQUENCE = 7;  // 一毫秒内支持生成128个id

    private static $sequence_id    = 0;
    private static $last_timestamp = 0;
    private static $server_id = 0;


    /**
     * 获取id
     * @param int $worker_id
     * @return float|int
     */
    public static function get($worker_id = 0)
    {
        //初始化id位
        $id = pow(2, 62);

        /* 1. 时间戳 41位 */
        $time = self::timeGen();
        $diff_time = $time - self::EPOCH_TIME;

        $shift = self::BITS_FULL - self::BITS_PRE - self::BITS_TIME;
        $id |= $diff_time << $shift;

        /* 2. 服务器id */
        $shift -= self::BITS_SERVER;
        $id |= (self::$server_id & (pow(2, self::BITS_SERVER) - 1)) << $shift;

        /* 3. 业务id */
        $shift -= self::BITS_WORKER;
        $id |= ($worker_id & (pow(2, self::BITS_WORKER) - 1)) << $shift;

        /* 4. 自增id */
        $id |= (self::$sequence_id % (pow(2, self::BITS_SEQUENCE) - 1));

        self::$sequence_id++;
        return $id;
    }

    /**
     * 获取当前时间
     */
    protected static function timeGen() {
        $wait_next_ms = 0;
        do {
            if($wait_next_ms > 0) {
                usleep(100); //等待下一毫秒，休眠0.1毫秒
            }

            $timestamp = microtime(true) * 1000;
            $timestamp = (int) $timestamp;

            if(self::$last_timestamp < $timestamp) {
                self::$sequence_id = 0;
            }

            $wait_next_ms++;

        } while (self::$last_timestamp == $timestamp
        && self::$sequence_id >= (pow(2, self::BITS_SEQUENCE) - 1)
        || self::$last_timestamp > $timestamp);

        self::$last_timestamp = $timestamp;
        return $timestamp;
    }
}