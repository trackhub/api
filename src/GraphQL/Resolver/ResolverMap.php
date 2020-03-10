<?php

namespace App\GraphQL\Resolver;

use App\Entity\Track;
use App\Entity\User\User;
use Overblog\GraphQLBundle\Definition\ArgumentInterface;

class ResolverMap extends \Overblog\GraphQLBundle\Resolver\ResolverMap
{
    protected function map()
    {
        return [
            'RootQuery' => [
                self::RESOLVE_FIELD => function ($value, ArgumentInterface $args, \ArrayObject $context, ResolveInfo $info) {
                    dump(func_get_args());
                    die;
                },
                'tracks' => function ($TODO, ArgumentInterface $ai, \ArrayObject $resolveInfo) {
                    return [
                        (function () {
                            $t = new Track(new User());
                            $t->setNameEn('track1');
                            return $t;
                        })(),
                        new Track(new User()),
                        new Track(new User()),
                    ];
                },


            ],

        ];
    }
}
