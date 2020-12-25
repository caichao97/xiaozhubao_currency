<?php


namespace AppBundle\Response;


use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Response;

class BaseResponse extends JsonResponse
{
    /**
     * @var integer
     */
    protected $code;

    /**
     * @var string
     */
    protected $message;

    /**
     * @var mixed
     */
    protected $dataContent;

    /**
     * BaseResponse constructor.
     * @param mixed $dataContent
     * @param string $message
     * @param int $code
     * @param int $status
     */
    public function __construct($dataContent = [], $message = 'success', $code = Response::HTTP_OK, $status = self::HTTP_OK)
    {
        $this->setDataContent($dataContent);
        $this->setMessage($message);
        $this->setCode($code);
        parent::__construct($this->getData(), $status, [], false);
    }

    /**
     * @return int
     */
    public function getCode(): int
    {
        return $this->code;
    }

    /**
     * @param int $code
     * @return BaseResponse
     */
    public function setCode(int $code): BaseResponse
    {
        $this->code = $code;
        return $this;
    }

    /**
     * @return string
     */
    public function getMessage(): string
    {
        return $this->message;
    }

    /**
     * @param string $message
     * @return BaseResponse
     */
    public function setMessage(string $message): BaseResponse
    {
        $this->message = $message;
        return $this;
    }

    /**
     * @return mixed
     */
    public function getDataContent()
    {
        return $this->dataContent;
    }

    /**
     * @param mixed $dataContent
     * @return BaseResponse
     */
    public function setDataContent($dataContent): BaseResponse
    {
        $this->dataContent = $dataContent;
        return $this;
    }


    /**
     * @return array
     */
    private function getData()
    {
        $data = [
            'status' => $this->getCode(),
            'content' => $this->getDataContent(),
            'timeStamp' => time(),
        ];
        if($this->code != 200){
            $data['errorMsg'] = $this->getMessage();
        }
        return $data;
    }
}