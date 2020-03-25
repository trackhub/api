<?php

namespace App\GraphQL\Resolver;

use App\Repository\PlaceRepository;

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
                    // @FIXME WIP
                    return $repo->findAll();
                }
            ]
        ];
    }
}
