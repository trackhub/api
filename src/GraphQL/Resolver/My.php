<?php


namespace App\GraphQL\Resolver;


use Overblog\GraphQLBundle\Definition\Resolver\ResolverInterface;

class My implements ResolverInterface
{
    public function sayHelloA($name) {
        die('test qwe123');
        throw new \Exception("Debug" . __METHOD__);
    }
}
