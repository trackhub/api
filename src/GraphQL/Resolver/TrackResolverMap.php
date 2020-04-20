<?php

declare(strict_types=1);

namespace App\GraphQL\Resolver;

use App\Repository\TrackRepository;
use Overblog\GraphQLBundle\Definition\ArgumentInterface;
use Overblog\GraphQLBundle\Resolver\ResolverMap;

class TrackResolverMap extends ResolverMap
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
                    // this endpoint is still in development!
                    if ($root !== null) {
                        throw new \Exception("not implemented");
                    }

                    $qb = $trackRepo->createQueryBuilder('t');
                    $qb->setMaxResults($ai['limit']);

                    $tracks = $qb->getQuery()->getResult();

                    return $tracks;
                },

                'tracksMap' => function($root, ArgumentInterface $ai, \ArrayObject $resolveInfo) use ($trackRepo) {
                    if ($root !== null) {
                        throw new \Exception("not implemented");
                    }

                    $qb = $trackRepo->createQueryBuilder('t');
                    $trackRepo->andWhereInCoordinates(
                        $qb,
                        $ai['skipTracks'],
                        $ai['neLat'],
                        $ai['swLat'],
                        $ai['neLon'],
                        $ai['swLon'],
                    );

                    $trackRepo->andWhereTrackIsPublic($qb);
                    $qb->setMaxResults(10);

                    return $trackRepo->findAll();
                }
            ],
        ];
    }
}
