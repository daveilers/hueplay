//
//  main.swift
//  asleep
//
//  Created by David Eilers on 12/4/17.
//  Copyright © 2017 David Eilers. All rights reserved.
//

import Foundation
import Quartz

let cfd = Quartz.CGSessionCopyCurrentDictionary()
if (cfd == nil) {
   exit(0)
}
if let d = cfd! as? [String: AnyObject] {
    if let locked = d["CGSSessionScreenIsLocked"] {
        if locked.boolValue {
            exit(0)
        }
    }
    
    let consoleKey = d["kCGSSessionOnConsoleKey"]
    if consoleKey == nil {
        exit(0)
    }
    if consoleKey! as! Bool {
        exit(1)
    }
    exit(0)
}
