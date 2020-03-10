<?php

namespace App\GraphQL\Resolver;

use App\Entity\Track;
use App\Entity\User\User;
use ArrayObject as ArrayObjectAlias;
use Overblog\GraphQLBundle\Definition\ArgumentInterface as ArgumentInterfaceAlias;

class ResolverMap extends \Overblog\GraphQLBundle\Resolver\ResolverMap
{
    protected function map()
    {
        return [
            'RootQuery' => [
//                self::RESOLVE_FIELD => function  ($value, ArgumentInterfaceAlias $args, ArrayObjectAlias $context, ResolveInfo $info) {
//                    return function() {
//                        var_dump(func_get_args());
//                    };
//                },
                'tracks' => function() {
//                    var_dump(func_get_args());
                        return [
                            new Track(new User()),
                        ];
                },
            ],
        ];
    }

    public function sayHelloA($name) {
//        die('test qwe123');
//        throw new \Exception("Debug" . __METHOD__);
    }

}
