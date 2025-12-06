export type Player = { 
    id: number; 
    clubId: number; 
    name: string; 
    position: string; 
    overall: number; 
    nationality: string; 
};

export type Club = { 
    id: number; 
    leagueId: number; 
    name: string;
    stadium: string;
};

export type League = { 
    id: number; 
    name: string; 
    country: string; 
};
