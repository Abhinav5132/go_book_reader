export namespace models {
	
	export class Library {
	    ID: number;
	    Name: string;
	    Path: string;
	    Picture: string;
	    Books: Book[];
	
	    static createFrom(source: any = {}) {
	        return new Library(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Path = source["Path"];
	        this.Picture = source["Picture"];
	        this.Books = this.convertValues(source["Books"], Book);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Franchise {
	    Id: number;
	    FranchiseName: string;
	    FranchiseImage: string;
	    Books: Book[];
	
	    static createFrom(source: any = {}) {
	        return new Franchise(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.FranchiseName = source["FranchiseName"];
	        this.FranchiseImage = source["FranchiseImage"];
	        this.Books = this.convertValues(source["Books"], Book);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Book {
	    Id: number;
	    Name: string;
	    Path: string;
	    FileType: string;
	    // Go type: time
	    LastAccessed: any;
	    Franchises: Franchise[];
	    LibraryID: number;
	    Library: Library;
	
	    static createFrom(source: any = {}) {
	        return new Book(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.Name = source["Name"];
	        this.Path = source["Path"];
	        this.FileType = source["FileType"];
	        this.LastAccessed = this.convertValues(source["LastAccessed"], null);
	        this.Franchises = this.convertValues(source["Franchises"], Franchise);
	        this.LibraryID = source["LibraryID"];
	        this.Library = this.convertValues(source["Library"], Library);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

