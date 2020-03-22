<?php

// this a debug file
// will be removed

namespace App\GraphQL\Resolver;

use GraphQL\Type\Definition\ResolveInfo;
use Overblog\GraphQLBundle\Definition\Resolver\ResolverInterface;

class MyResolver implements ResolverInterface
{
    public function test($entity, ResolveInfo $resolveInfo)
    {
        if ($resolveInfo->fieldName === 'order') {
            return $entity->getOrder();
        }
    }
}
