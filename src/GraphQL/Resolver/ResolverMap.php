<?php

namespace App\GraphQL\Resolver;

use App\Entity\Track;
use App\Entity\User\User;
use Overblog\GraphQLBundle\Definition\ArgumentInterface;

class ResolverMap extends \Overblog\GraphQLBundle\Resolver\ResolverMap
{
    protected function map()
    {
        // test for DI
        $prefix = 'test 123';

        return [
            'RootQuery' => [
                self::RESOLVE_FIELD => function ($value, ArgumentInterface $args, \ArrayObject $context, ResolveInfo $info) {
                    dump(func_get_args());
                    die;
                },
                'tracks' => function ($root, ArgumentInterface $ai, \ArrayObject $resolveInfo) use ($prefix) {
                    if ($root !== null) {
                        throw new \Exception("not implemented");
                    }

                    $limit = $ai['limit'];

                    $data = [
                        (function ($prefix) {
                            $t = new Track(new User());
                            $t->setNameEn($prefix . ' Markovo');
                            return $t;
                        })($prefix),
                        (function () {
                            $t = new Track(new User());
                            $t->setNameEn('Boikovo');
                            return $t;
                        })(),
                        new Track(new User()),
                    ];

                    return array_slice($data, 0, $limit);
                },
            ],
        ];
    }
}
