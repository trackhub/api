<?php

namespace App\GraphQL\Resolver;

use App\Entity\Track;
use GraphQL\Type\Definition\ResolveInfo;
use Overblog\GraphQLBundle\Definition\Resolver\ResolverInterface;

class My implements ResolverInterface
{
    public function sayHelloA(Track $track, ResolveInfo $resolveInfo)
    {
        if ($resolveInfo->fieldName === 'nameEn') {
            return $track->{'get' . $resolveInfo->fieldName}() . ' from ' . __METHOD__;
        }

        return $track->{'get' . $resolveInfo->fieldName}();
    }
}
