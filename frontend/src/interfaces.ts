export interface Tea {
    teaName: string;
    blendDescription: string;
    brewingDuration: number;
    color: string;
    containerWeight: number;
    id: string;
    inUse: number;
    initialWeight: number;
    origin: Origin
    portionWeight: number;
    size: string;
    teaType: string;
    temperature: number;
}

export interface Origin {
   shopName: string;
   shopLocation: string;
}