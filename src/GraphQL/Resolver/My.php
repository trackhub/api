<?php


namespace App\GraphQL\Resolver;


use Overblog\GraphQLBundle\Definition\Resolver\ResolverInterface;

class My implements ResolverInterface
{
    public function sayHelloA() {
        return uniqid(__METHOD__);
//        throw new \Exception("Debug" . __METHOD__);
    }
}
