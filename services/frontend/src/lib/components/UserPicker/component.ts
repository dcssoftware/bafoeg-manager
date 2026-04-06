// Using null as default to represent "no user selected" externally
// Internally, we convert this to a special string value for Bits UI compatibility
export const defaultUserID: string | null = null;

export function defaultEventChangeSelectedUserID(userID: string | null) {
  // console.log("handler", { userID });
}
