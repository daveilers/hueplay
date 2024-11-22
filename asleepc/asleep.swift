//
//  main.swift
//  asleep
//
//  Created by David Eilers on 12/4/17.
//  Copyright © 2017 David Eilers. All rights reserved.
//
// exits 0 if asleep (locked or no current logged in user), 1 if awake

import Foundation
import Quartz

@_cdecl("Asleep")
public func asleep() -> Bool {
    let cfd = Quartz.CGSessionCopyCurrentDictionary()
    if cfd == nil {
        return true
    }
    if let d = cfd! as? [String: AnyObject] {
        if let locked = d["CGSSessionScreenIsLocked"] {
            if locked.boolValue {
                return true
            }
        }
        let consoleKey = d["kCGSSessionOnConsoleKey"]
        if consoleKey == nil {
            return true
        }
        if consoleKey! as! Bool {
            return false
        }
        return true
    }
    return true
}
