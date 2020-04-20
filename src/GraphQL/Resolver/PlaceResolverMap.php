<?php

namespace App\GraphQL\Resolver;

use App\Repository\PlaceRepository;
use Overblog\GraphQLBundle\Definition\ArgumentInterface;

class PlaceResolverMap extends \Overblog\GraphQLBundle\Resolver\ResolverMap
{
    private PlaceRepository $placeRepo;

    public function __construct(PlaceRepository $placeRepo)
    {
        $this->placeRepo = $placeRepo;
    }

    protected function map()
    {
        $repo = $this->placeRepo;

        return [
            'RootQuery' => [
                'places' => function() use ($repo) {
                    // this endpoint is still in development!

                    return $repo->findAll();
                },
                'placesMap' => function($root, ArgumentInterface $ai, \ArrayObject $resolveInfo) use ($repo) {
                    $qb = $repo->createQueryBuilder('p');
                    $qb->setMaxResults(10);
                    $repo->andWhereInCoordinates(
                        $qb,
                        $ai['skipPlaces'],
                        $ai['neLat'],
                        $ai['swLat'],
                        $ai['neLon'],
                        $ai['swLon'],
                    );

                    return $qb->getQuery()->getResult();
                },
            ]
        ];
    }
}
