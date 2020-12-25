<?php

namespace AppBundle\Controller;

use AppBundle\Response\BaseResponse;
use AppBundle\Services\IDService;
use Symfony\Component\Routing\Annotation\Route;

/**
 * Class IDController
 * @package AppBundle\Controller
 * @Route("ID")
 */
class IDController extends BaseController
{

    /**
     * 获取全局唯一ID
     * @Route("/getId")
     */
    public function getIdAction(){
       return new BaseResponse(IDService::get());
    }
}
