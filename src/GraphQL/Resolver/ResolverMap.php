<?php

namespace App\GraphQL\Resolver;

use App\Repository\TrackRepository;
use Overblog\GraphQLBundle\Definition\ArgumentInterface;

class ResolverMap extends \Overblog\GraphQLBundle\Resolver\ResolverMap
{
    private TrackRepository $trackRepo;

    public function __construct(TrackRepository $trackRepo)
    {
        $this->trackRepo = $trackRepo;
    }

    protected function map()
    {
        $trackRepo = $this->trackRepo;

        return [
            'RootQuery' => [
                'tracks' => function ($root, ArgumentInterface $ai, \ArrayObject $resolveInfo) use ($trackRepo) {
                    if ($root !== null) {
                        throw new \Exception("not implemented");
                    }

                    $limit = $ai['limit'];

                    $tracks = $trackRepo->findAll();

                    return $tracks;
                },

                'tracksMap' => function() use ($trackRepo) {
                    return $trackRepo->findAll();
                }
            ],
        ];
    }
}
